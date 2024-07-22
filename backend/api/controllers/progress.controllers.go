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

type ProgressControllers struct {
	Services *services.ProgressServices
}

func NewProgressControllers() *ProgressControllers {
	return &ProgressControllers{
		Services: services.NewProgressServices(),
	}
}

func (pc *ProgressControllers) Create(c *gin.Context) {
	// get penelitiID

	projectID, _ := strconv.Atoi(c.Param("id"))

	input := new(dto.ProgressReport)
	if err := c.ShouldBindJSON(&input); err != nil {
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

	if err := pc.Services.CreateReport(uint(projectID), *input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "mengunggah laporan berhasil")
}
