package middleware

import (
	"net/http"
	"sync"

	"strconv"
	"time"

	"github.com/bagasa11/banturiset/api/services"
	"github.com/gin-gonic/gin"
)

type tahap struct {
	Tahap uint8 `json:"tahap"`
}

type tahapInputValidate struct {
	tahap     uint8
	projectID uint
}

var inputChannel chan tahapInputValidate

func sendToChan(c chan<- tahapInputValidate, input tahapInputValidate) {
	c <- input
}

func handleChan(ch <-chan tahapInputValidate, wg *sync.WaitGroup, result *tahapInputValidate) {
	defer wg.Done()
	for v := range ch {
		result = &v
		// result = &v
		return
	}

}

func DateAndStageValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "url tidak valid")
			c.Abort()
			return
		}
		penelitiID, exist := c.Get("role_id")
		if !exist {
			c.JSON(400, "header id peneliti tidak ditemukan")
			c.Abort()
			return
		}

		if err := services.IsMyProject(uint(projectID), penelitiID.(uint)); err != nil {
			c.JSON(http.StatusForbidden, "tidak boleh mengedit proyek milik user lain")
			c.Abort()
			return
		}
		// mendapatkan data tahapan berdasarkan tahap dan project id
		tahapInput := new(tahap)
		if err := c.ShouldBindJSON(&tahapInput); err != nil {
			c.JSON(400, "gagal mengambil data input")
			c.Abort()
			return
		}
		dataTahap, err := services.GetDataByTahap(uint(projectID), tahapInput.Tahap)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		// melakukan komparasi range data dengan now()
		now := time.Now()
		if !(now.After(dataTahap.Start) && now.Before(dataTahap.End)) {
			c.JSON(http.StatusForbidden, "waktu pelaksanaan kegiatan belum dimulai")
			c.Abort()
			return
		}
		c.Next()
	}
}

func EnsureProjectWasClosed() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "url tidak valid")
			c.Abort()
			return
		}

		if services.IsOpenFund(uint(projectID)) != nil {
			c.JSON(http.StatusForbidden, "tidak dapat membuat laporan saat waktu pendanaan masih dibuka")
			c.Abort()
			return
		}
		c.Next()
	}
}

func ValidateCreateReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "url tidak valid")
			c.Abort()
			return
		}
		penelitiID, exist := c.Get("role_id")
		if !exist {
			c.JSON(400, "header id peneliti tidak ditemukan")
			c.Abort()
			return
		}
		tahapInput := new(tahap)
		if err := c.ShouldBindJSON(&tahapInput); err != nil {
			c.JSON(400, "gagal mengambil data input")
			c.Abort()
			return
		}

		project, err := services.MyProjectWasClosedDetail(uint(projectID), penelitiID.(uint), tahapInput.Tahap)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		// length check
		if len(project.Tahapan) <= 0 {
			c.JSON(http.StatusNotFound, "data tahapan tidak ditemukan")
			c.Abort()
			return
		}

		t := time.Now()
		if project.FundUntil.After(t) {
			c.JSON(http.StatusForbidden, "tidak dapat membuat laporan saat waktu pendanaan masih dibuka")
			c.Abort()
			return
		}

		if !(project.Tahapan[0].Start.Before(t) && project.Tahapan[0].End.After(t)) {
			c.JSON(http.StatusForbidden, "waktu pelaksanaan kegiatan belum dimulai")
			c.Abort()
			return
		}
		// mychan := make(chan tahap)
		// mychan <- *tahapInput
		sendToChan(inputChannel, tahapInputValidate{tahap: tahapInput.Tahap, projectID: uint(projectID)})
		c.Next()
	}
}

func EnsureNotRedundant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var wg sync.WaitGroup
		wg.Add(1)

		data := new(tahapInputValidate)
		go handleChan(inputChannel, &wg, data)
		wg.Wait()

		if data == nil {
			ctx.JSON(500, "terjadi kesalahan pada goroutine")
			ctx.Abort()
			return
		}
		ps := services.NewProgressServices()
		if err := ps.IsRedundant(data.projectID, data.tahap); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}
