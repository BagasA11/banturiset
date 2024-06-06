package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bagasa11/banturiset/api/dto"

	"github.com/bagasa11/banturiset/api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PengajuanControllers struct {
	Service *services.PengajuanServices
}

func NewPengajuanController() *PengajuanControllers {
	return &PengajuanControllers{
		Service: services.NewPengajuanService(),
	}
}

func (pc *PengajuanControllers) Create(c *gin.Context) {

	// header validation
	role, exist := c.Get("role")
	if !exist {
		c.JSON(http.StatusBadRequest, "header role tidak ditemukan")
		return
	}

	if strings.ToLower(role.(string)) != "penyunting" {
		c.JSON(http.StatusForbidden, "laman khusus admin")
		return
	}

	roleID, exist := c.Get("role_id")
	if !exist {
		c.JSON(http.StatusBadRequest, "header role_id tidak ditemukan")
		return
	}

	req := new(dto.Pengajuan)
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

	if err := pc.Service.Create(*req, roleID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal menambah data pengajuan",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "sukses")
}

func (pc *PengajuanControllers) Open(c *gin.Context) {
	p, err := pc.Service.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": p,
		"len":  len(p),
	})
}
