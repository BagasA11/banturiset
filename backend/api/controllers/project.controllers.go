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

type ProjectControllers struct {
	Service *services.ProjectService
}

func NewProjectControllers() *ProjectControllers {
	return &ProjectControllers{
		Service: services.NewProjectService(),
	}
}

func (pc *ProjectControllers) Create(c *gin.Context) {
	// role validation
	if role, _ := c.Get("role"); strings.ToLower(role.(string)) != "peneliti" {
		c.JSON(http.StatusUnauthorized, "laman khusus peneliti")
		return
	}

	role_id, exist := c.Get("role_id")
	if !exist {
		c.JSON(http.StatusBadRequest, "detail user tidak ditemukan")
		return
	}

	req := new(dto.CreateProject)
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

	if err := pc.Service.Create(*req, role_id.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal membuat project",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "sukses")
}

func (pc *ProjectControllers) MyProject(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// role validation
	if role, _ := c.Get("role"); strings.ToLower(role.(string)) != "peneliti" {
		c.JSON(http.StatusUnauthorized, "laman khusus peneliti")
		return
	}

	role_id, exist := c.Get("role_id")
	if !exist {
		c.JSON(http.StatusBadRequest, "detail user tidak ditemukan")
		return
	}

	projects, err := pc.Service.MyProject(role_id.(uint), uint(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal mengambil list proyek",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": projects,
	})
}

func (pc *ProjectControllers) UploadProposal(c *gin.Context) {
	// role validation
	if role, _ := c.Get("role"); strings.ToLower(role.(string)) != "peneliti" {
		c.JSON(http.StatusUnauthorized, "laman khusus peneliti")
		return
	}

	role_id, exist := c.Get("role_id")
	if !exist {
		c.JSON(http.StatusBadRequest, "detail user tidak ditemukan")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "project tidak ditemukan atau format id salah",
			"error": err.Error(),
		})
		return
	}

	req := new(dto.Proposal)
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

	if err := pc.Service.UploadProposal(uint(id), role_id.(uint), req.Url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "gagal mengunggah proposal",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (pc *ProjectControllers) UploadKlirens(c *gin.Context) {
	// role validation
	if role, _ := c.Get("role"); strings.ToLower(role.(string)) != "peneliti" {
		c.JSON(http.StatusUnauthorized, "laman khusus peneliti")
		return
	}

	role_id, exist := c.Get("role_id")
	if !exist {
		c.JSON(http.StatusBadRequest, "detail user tidak ditemukan")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "project tidak ditemukan atau format id salah",
			"error": err.Error(),
		})
		return
	}

	req := new(dto.Klirens)
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

	if err := pc.Service.UploadProposal(uint(id), role_id.(uint), req.Url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "gagal mengunggah proposal",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (pc *ProjectControllers) Reject(c *gin.Context) {
	// role validation
	role, exist := c.Get("role")
	if !exist {
		c.JSON(http.StatusBadRequest, "header role diperlukan")
		return
	}

	if strings.ToLower(role.(string)) != "penyunting" {
		c.JSON(http.StatusForbidden, "laman khusus admin")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"pesan": "project tidak ditemukan atau format id salah",
			"error": err.Error(),
		})
		return
	}

	req := new(dto.ProjectDitolak)
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

	if err := pc.Service.Tolak(uint(id), *req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal mereject proyek",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "sukses")
}
