package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	tz "github.com/bagasa11/banturiset/timezone"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastseen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*client)
)

func PerClientRateLimiter() gin.HandlerFunc {
	go func() {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastseen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
	}()

	return func(ctx *gin.Context) {
		ip, _, err := net.SplitHostPort(ctx.Request.RemoteAddr)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
		}
		clients[ip].lastseen = tz.GetTime(time.Now())
		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			ctx.JSON(http.StatusTooManyRequests, "api reached capacity")
			ctx.Abort()
			return
		}

		mu.Unlock()
		ctx.Next()
	}
}

func SimpleLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter := rate.NewLimiter(2, 4)
		if !limiter.Allow() {
			ctx.JSON(http.StatusTooManyRequests, "api limit is reached")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
