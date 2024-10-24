package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/services"
	erf "github.com/bagasa11/banturiset/errorf"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProgressControllers struct {
	Services *services.ProgressServices
}

func NewProgressControllers() *ProgressControllers {
	return &ProgressControllers{
		Services: services.NewProgressServices(),
	}
}

func (pc *ProgressControllers) CreateClean(c *gin.Context) {
	// get penelitiID
	penelitiID, exist := c.Get("role_id")
	if !exist {
		c.JSON(400, "header id peneliti tidak ditemukan")
		c.Abort()
		return
	}
	// project id
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, "url tidak valid")
		c.Abort()
		return
	}

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

	// // validate project state
	// project, err := services.MyProjectWasClosedDetail(uint(projectID), penelitiID.(uint), input.Tahap)
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"error": "data project ditemukan",
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// // length check
	// if len(project.Tahapan) <= 0 {
	// 	c.JSON(http.StatusNotFound, "data tahapan tidak ditemukan")
	// 	return
	// }

	// t := tz.GetTime(time.Now())

	// if project.FundUntil.After(t) {
	// 	c.JSON(http.StatusForbidden, "tidak dapat membuat laporan saat waktu pendanaan masih dibuka")
	// 	c.Abort()
	// 	return
	// }

	// if !(project.Tahapan[0].Start.Before(t) && project.Tahapan[0].End.After(t)) {
	// 	c.JSON(http.StatusForbidden, "waktu pelaksanaan kegiatan belum dimulai")
	// 	return
	// }

	err = services.ClosedProjectChecker(uint(projectID), penelitiID.(uint), input.Tahap)
	status, msg := handleErrorMessage(err)

	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
			"pesan": msg,
		})
		return
	}

	// redundant progress check
	if err := pc.Services.IsRedundantProgress(uint(projectID), input.Tahap); err != nil {
		if err == erf.ErrRedundant {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": fmt.Sprintf("data laporan tahap ke-%d redundan", input.Tahap),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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

func handleErrorMessage(err error) (statusCode int, message string) {

	if err == gorm.ErrRecordNotFound {
		return 404, "data project tidak ditemukan"
	}
	if err == erf.ErrNilTahap {
		return http.StatusUnprocessableEntity, "tahap tidak boleh kosong"
	}

	if err == erf.ErrDonationStillOpen {
		return http.StatusUnprocessableEntity, "tidak dapat membuat laporan karena donasi masih dibuka"
	}

	if err == erf.ErrHaveNotStartEvent {
		return http.StatusForbidden, "waktu pelaksanaan tahapan belum dimulai"
	}
	return 200, ""
}
