package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bagasa11/banturiset/api/dto"

	"github.com/bagasa11/banturiset/api/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Donasi struct {
	Service *services.DonasiService
}

func NewDonasi() *Donasi {
	return &Donasi{
		Service: services.NewDonasiService(),
	}
}

func (dc *Donasi) CreateDonasi(c *gin.Context) {

	roleID, _ := c.Get("role_id")
	if roleID.(uint) == 0 {
		c.JSON(http.StatusBadRequest, "id donatur diperlukan")
		return
	}

	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "parameter id diperlukan",
			"error": err.Error(),
		})
		return
	}

	email, _ := c.Get("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, "header email diperlukan")
	}

	req := new(dto.CreateDonasi)
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

	m, err := dc.Service.Create(roleID.(uint), uint(projectID), *req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	page, err := dc.Service.CreateInvoice(m, email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": page,
	})
}

func (dc *Donasi) GetInvoiceDetail(c *gin.Context) {

	roleID, _ := c.Get("role_id")
	if roleID.(uint) == 0 {
		c.JSON(http.StatusBadRequest, "id donatur diperlukan")
		return
	}

	tr := c.Param("id")
	if tr == "" {
		c.JSON(http.StatusBadRequest, "id transaksi diperlukan")
		return
	}

	data, err := dc.Service.GetTransaction(tr, roleID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (dc *Donasi) Notif(c *gin.Context) {
	req := new(dto.NotifInvoice)
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

	if c.GetHeader("x-callback-token") != os.Getenv("XENDIT_CALLBACK") {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "invalid callback token",
		})
		return
	}

	if err := dc.Service.Notifikasi(*req); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
