package controllers

import (
	"fmt"
	"net/http"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/services"
	"github.com/bagasa11/banturiset/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	Auth *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		Auth: services.NewAuthService(),
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	req := new(dto.Login)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "Invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	if req == nil {
		c.JSON(http.StatusUnprocessableEntity, "body request kosong")
		return
	}

	token, err := ac.Auth.Login(*req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"pesan": "autentikasi invalid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan": "login sukses",
		"token": token,
	})
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	req := new(dto.RefreshTokenRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "Invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	newToken, err := helpers.UpdateToken(req.OldToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan": "update token sukses",
		"token": newToken,
	})
}
