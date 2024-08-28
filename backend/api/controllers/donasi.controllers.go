package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	userID, _ := c.Get("id")
	if userID.(uint) == 0 {
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

	m, err := dc.Service.Create(userID.(uint), uint(projectID), *req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	page, err := dc.Service.CreateInvoice(m, email.(string))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": page,
	})
}

func (dc *Donasi) GetInvoiceDetail(c *gin.Context) {

	roleID, _ := c.Get("id")
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

	if strings.ToLower(req.Status) != "paid" {

		if err := dc.Service.UpdateStatus(req.ExternalID, req.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, "ok")
		return
	}

	err := dc.Service.ConfirmPayment(req.ExternalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (dc *Donasi) GetHistory(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "parameter id invalid")
		return
	}

	donasi, err := dc.Service.GetAllHistory(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal mendapatkan data histori donasi",
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"data":   donasi,
		"length": len(donasi),
	})
}

func (dc *Donasi) Contributors(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "parameter limit invalid")
		return
	}

	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id project invalid")
		return
	}
	contribrutors, err := dc.Service.Contributors(uint(projectID), uint(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"pesan": "gagal mengambil data contributor",
		})
		return
	}
	c.JSON(200, gin.H{
		"data":   contribrutors,
		"length": len(contribrutors),
	})
}

func (dc *Donasi) MyHistory(c *gin.Context) {
	donaturID, _ := c.Get("role_id")
	if donaturID == nil {
		c.JSON(http.StatusBadRequest, "donatur ID tidak ditemukan")
		return
	}

	d, err := dc.Service.MyHistory(donaturID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": d,
	})
}
