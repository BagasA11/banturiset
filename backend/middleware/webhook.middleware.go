package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bagasa11/banturiset/config"
	"github.com/gin-gonic/gin"
)

func CheckRedundantWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {

		webhook_id := c.GetHeader("webhook-id")
		fmt.Println(webhook_id)
		if webhook_id == "" {
			c.JSON(http.StatusBadRequest, "webhook-id tidak ditemukan")
			c.Abort()
			return
		}

		// memeriksa apakah ada data dengan cache key yang sama
		// jika tidak ditemukan, maka akan melakukan setup data
		// namun jika ditemukan, maka akan skip

		cache := config.GetCacheTTL()
		_, err := cache.Get(webhook_id)

		if err == nil {
			c.AbortWithStatus(200)
			return
		}

		// if key not found
		if err == config.CacheNotFound {
			fmt.Println("set cache with a key: ", webhook_id)
			cache.SetWithTTL(webhook_id, webhook_id, 45*time.Minute)
		}

		c.Next()
	}
}

func VerifyWebhookToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("x-callback-token") != os.Getenv("XENDIT_WEBHOOK") {
			fmt.Println("callback token invalid")
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": "invalid callback token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
