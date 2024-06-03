package controllers

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/services"
	"github.com/bagasa11/banturiset/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UsersController struct {
	Services *services.UserService
}

func NewUsersController() *UsersController {
	return &UsersController{
		Services: services.NewUserService(),
	}
}

func (uc *UsersController) UserRegistration(c *gin.Context) {
	req := new(dto.UserRegister)
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

	if !helpers.ValidatePattern("email", req.Email) {
		c.JSON(http.StatusUnprocessableEntity, "format email ditolak")
		return
	}

	if !helpers.ValidatePattern("phone", req.Phone) {
		c.JSON(http.StatusUnprocessableEntity, "format handphone ditolak")
		return
	}

	if !helpers.ValidatePattern("post", req.PostCode) {
		c.JSON(http.StatusUnprocessableEntity, "format kode pos ditolak")
		return
	}

	if strings.Contains(strings.ToLower(req.Role), "admin") {
		c.JSON(http.StatusUnauthorized, "admin tidak boleh melakukan registrasi!!!")
		return
	}

	if !slices.Contains([]string{"peneliti", "donatur"}, strings.ToLower(req.Role)) {
		c.JSON(http.StatusUnauthorized, "pilih opsi peneliti atau donatur")
		return
	}

	userID, err := uc.Services.UserRegister(*req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal menginput data user",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"pesan":   "simpan user_id untuk proses registrasi lebih lanjut. Masukkan user_id ke dalam parameter url",
	})
}

func (uc *UsersController) DonaturCreate(c *gin.Context) {
	// example.com/api/user/create/1?role=

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "tidak dapat mengonversi id user menjadi integer",
			"error": err.Error(),
		})
		return
	}

	req := new(dto.DonaturRegister)
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

	if strings.ToLower(req.Role) != "donatur" {
		c.JSON(http.StatusUnprocessableEntity, "diharapkan untuk menjadi donatur")
		return
	}

	if err := uc.Services.CheckID(uint(userID), req.Role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "user id tidak ditemukan atau peran tidak valid",
			"error": err.Error(),
		})
		return
	}

	if err := uc.Services.CreateDonatur(uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "tidak dapat menambahkan data donatur",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "registrasi sukses. Mohon tunggu verifikasi akun dari admin")
}

func (uc *UsersController) PenelitiCreate(c *gin.Context) {
	// example.com/api/user/create/1?role=

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "tidak dapat mengonversi id user menjadi integer",
			"error": err.Error(),
		})
		return
	}

	req := new(dto.PenelitiRegister)
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

	if strings.ToLower(req.Role) != "peneliti" {
		c.JSON(http.StatusUnprocessableEntity, "peran invalid")
		return
	}

	if err := uc.Services.CheckID(uint(userID), req.Role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "user id tidak ditemukan",
			"error": err.Error(),
		})
		return
	}

	if err := uc.Services.CreatePeneliti(uint(userID), *req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "tidak dapat menambahkan data peneliti",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "registrasi sukses. Mohon tunggu verifikasi akun dari admin")
}

func (uc *UsersController) NeedVerify(c *gin.Context) {
	role, exist := c.Get("role")
	if !exist {
		c.JSON(http.StatusBadRequest, "header tidak lengkap")
		return
	}
	if strings.ToLower(role.(string)) != "penyunting" || strings.ToLower(role.(string)) != "reviewer" {
		c.JSON(http.StatusForbidden, "laman khusus admin")
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "gagal mengambil parameter page",
			"error": err.Error(),
		})
		return
	}
	users, err := uc.Services.NotVerified(uint(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "tidak bisa mengambil info user",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
