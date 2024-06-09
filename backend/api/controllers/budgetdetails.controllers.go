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

type BudgetDetailsController struct {
	Service *services.BudgetDetailService
}

func NewBudgetDetailsController() *BudgetDetailsController {
	return &BudgetDetailsController{
		Service: services.NewBudgetDetailService(),
	}
}

func (bdc *BudgetDetailsController) Create(c *gin.Context) {
	role, exist := c.Get("role")
	if !exist {
		c.JSON(http.StatusBadRequest, "header role diperlukan")
		return
	}
	if strings.ToLower(role.(string)) != "peneliti" {
		c.JSON(http.StatusForbidden, "laman khusus peneliti")
	}

	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "parameter id diperlukan",
			"error": err.Error(),
		})
		return
	}

	req := new(dto.BudgetDetailsCreate)
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

	if err := bdc.Service.Create(uint(projectID), *req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal menambahkan data detail budget",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, "ok")
}

func (bc *BudgetDetailsController) Updates(c *gin.Context) {
	role, _ := c.Get("role")

	if strings.ToLower(role.(string)) != "peneliti" {
		c.JSON(http.StatusForbidden, "laman khusus peneliti")
	}

	// roleID
	roleID, exist := c.Get("role_id")
	if !exist {
		c.JSON(http.StatusBadRequest, "header role_id diperlukan")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "parameter id diperlukan",
			"error": err.Error(),
		})
		return
	}

	req := new(dto.BudgetDetailsCreate)
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

	if err := bc.Service.Updates(uint(id), *req, roleID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal mengupdate detail budget",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, "ok")
}

func (bc *BudgetDetailsController) Delete(c *gin.Context) {
	role, _ := c.Get("role")
	if strings.ToLower(role.(string)) != "peneliti" {
		c.JSON(http.StatusForbidden, "laman khusus peneliti")
	}

	// role id
	roleID, exist := c.Get("role_id")
	if !exist {
		c.JSON(400, "header role_id diperlukan")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "parameter id diperlukan",
			"error": err.Error(),
		})
		return
	}

	if err := bc.Service.Delete(uint(id), roleID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal menghapus detail budget",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, "ok")
}
