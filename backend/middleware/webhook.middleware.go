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

		t := time.Now()
		webhook_id := c.GetHeader("webhook-id")
		fmt.Println(webhook_id)
		if webhook_id == "" {
			c.JSON(http.StatusBadRequest, "webhook-id tidak ditemukan")
			c.Abort()
			return
		}

		// memeriksa apakah ada data dengan cache key yang sama
		// jika tidak ditemukan, maka akan melakukan setup data
		// namun jika ditemukan, akan memeriksa expired time dari cache
		// jika expired akan dihapus
		cache := config.GetCacheTTL()
		_, chDura, err := cache.GetWithTTL(webhook_id)

		if err == nil {
			// if expired then: remove item

			if time.Since(t) > chDura {
				err = cache.Remove(webhook_id)
				fmt.Println("remove cache item with key: ", webhook_id)
				fmt.Println("error when remove cache: ", err.Error())
			}
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
