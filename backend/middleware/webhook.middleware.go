package middleware

import (
	"net/http"
	"time"

	"github.com/bagasa11/banturiset/config"
	"github.com/gin-gonic/gin"
)

func CheckRedundantRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		webhook_id, _ := c.Get("webhook-id")
		if webhook_id.(string) == "" {
			c.JSON(http.StatusBadRequest, "webhook-id tidak ditemukan")
			c.Abort()
			return
		}
		// memeriksa apakah ada data dengan cache key yang sama
		// jika data ditemukan, mak akan melakukan abort
		cache := config.GetCacheTTL()
		_, chDura, err := cache.GetWithTTL(webhook_id.(string))
		if err == nil {
			// if expired then: remove item
			if time.Since(t) > chDura {
				cache.Remove(webhook_id.(string))
			}
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
