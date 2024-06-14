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
			c.JSON(http.StatusUnauthorized, "laman khusus peneliti")
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
			c.JSON(http.StatusBadRequest, "header role diperlukan")
			return
		}

		if !slices.Contains([]string{"penyunting", "admin", "penelaah"}, strings.ToLower(role.(string))) {
			c.JSON(http.StatusForbidden, "laman khusus admin")
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
			return
		}
		c.Next()
	}
}
