package middleware

import (
	"net/http"
	"strings"

	"github.com/bagasa11/banturiset/helpers"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, "Authorization header needed!")
			c.Abort()
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusBadRequest, "tipe token invalid")
			c.Abort()
			return
		}

		// get jwt Claims
		claims, err := helpers.ValidateToken(authHeaderParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
				"pesan": "token invalid",
			})
			c.Abort()
			return
		}

		// set header
		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("role_id", claims.RoleID)

		c.Next()
	}
}
