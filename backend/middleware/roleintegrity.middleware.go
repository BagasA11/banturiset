package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func PenelitiOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if role, _ := c.Get("role"); strings.ToLower(role.(string)) != "peneliti" {
			c.JSON(http.StatusForbidden, "laman khusus peneliti")
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// role validation
		role, exist := c.Get("role")
		if !exist {
			c.JSON(http.StatusForbidden, "header role diperlukan")
			c.Abort()
			return
		}

		if !slices.Contains([]string{"penyunting", "admin", "penelaah"}, strings.ToLower(role.(string))) {
			c.JSON(http.StatusForbidden, "laman khusus admin")
			c.Abort()
			return
		}

		c.Next()
	}
}

func DonaturOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if strings.ToLower(role.(string)) != "donatur" {
			c.JSON(http.StatusForbidden, "laman khusus donatur")
			c.Abort()
			return
		}
		c.Next()
	}
}
