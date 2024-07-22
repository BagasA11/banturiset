package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bagasa11/banturiset/api/dto"

	"github.com/bagasa11/banturiset/api/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TahapControllers struct {
	Service *services.TahapService
}

func NewTahapControllers() *TahapControllers {
	return &TahapControllers{
		Service: services.NewTahapService(),
	}
}

func (tc *TahapControllers) Create(c *gin.Context) {

	roleID, _ := c.Get("role_id")
	if roleID.(uint) == 0 {
		c.JSON(http.StatusBadRequest, "id peneliti diperlukan")
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

	req := new(dto.TahapCreate)
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

	if err := services.IsTahapRedundant(uint(projectID), req.Tahap); err != nil {
		if err.Error() == "data redundan" {
			c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("data kolom tahap %d redundan", req.Tahap))
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	if err := tc.Service.Create(uint(projectID), roleID.(uint), *req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal menambahkan data tahapan penelitian",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "request sukses")
}

func (tc *TahapControllers) List(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "parameter id diperlukan",
			"error": err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "parameter limit diperlukan",
			"error": err.Error(),
		})
		return
	}

	tahap, err := tc.Service.List(uint(projectID), uint(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "galat!!!",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tahap,
	})
}

func (tc *TahapControllers) Update(c *gin.Context) {

	roleID, _ := c.Get("role_id")
	if roleID.(uint) == 0 {
		c.JSON(http.StatusBadRequest, "role_id diperlukan")
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

	tahapID, err := strconv.Atoi(c.Param("tahapid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "parameter id diperlukan",
			"error": err.Error(),
		})
		return
	}

	req := new(dto.TahapCreate)
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

	if err := tc.Service.Update(uint(tahapID), *req, uint(projectID), roleID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal menambahkan data tahapan penelitian",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "request sukses")
}

func (tc *TahapControllers) Delete(c *gin.Context) {

	roleID, _ := c.Get("role_id")
	if roleID.(uint) == 0 {
		c.JSON(http.StatusBadRequest, "role_id diperlukan")
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

	tahapID, err := strconv.Atoi(c.Param("tahapid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "parameter id diperlukan",
			"error": err.Error(),
		})
		return
	}

	if err := tc.Service.Delete(uint(tahapID), uint(projectID), roleID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal menghapus data",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
