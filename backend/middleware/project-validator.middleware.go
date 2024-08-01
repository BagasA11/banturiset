package middleware

import (
	"net/http"

	"strconv"

	_ "github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/services"
	"github.com/gin-gonic/gin"
)

// type PengajuanID struct {
// 	ID uint `json:"pengajuan_id" binding:"required"`
// }

// func SkemaValidation() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		pengajuan_id := new(dto.CreateProject)
// 		if err := ctx.ShouldBindJSON(&pengajuan_id); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			ctx.Abort()
// 			return
// 		}

// 		ps := services.NewPengajuanService()
// 		err := ps.IsOpen(pengajuan_id.PengajuanID)
// 		if err == gorm.ErrRecordNotFound {
// 			ctx.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("skema penelitian %d sudah ditutup", pengajuan_id.PengajuanID))
// 			ctx.Abort()
// 			return
// 		}
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			ctx.Abort()
// 			return
// 		}

// 		ctx.Next()
// 	}
// }

func SubmitValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"pesan": "format url invalid",
			})
			c.Abort()
			return
		}

		penelitiID, exist := c.Get("role_id")
		if !exist {
			c.JSON(http.StatusBadRequest, "id peneliti tidak ditemukan")
			c.Abort()
			return
		}

		ps := services.NewProjectService()
		project, err := ps.Preview(uint(projectID), penelitiID.(uint))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
				"pesan": "konten tidak ditemukan",
			})
			c.Abort()
			return
		}

		if project.ProposalUrl == nil {
			c.JSON(http.StatusUnprocessableEntity, "proposal penelitian tidak boleh kosong")
			c.Abort()
			return
		}

		if len(project.Tahapan) <= 0 {
			c.JSON(http.StatusUnprocessableEntity, "detail tahap pelaksanaan project tidak boleh kosong")
			c.Abort()
			return
		}

		if uint8(project.Milestone) != project.Tahapan[0].Tahap {
			c.JSON(http.StatusUnprocessableEntity, "tahap pelaksanaan project dan jumlah milestone project tidak sama")
			c.Abort()
			return
		}

		if len(project.BudgetDetails) <= 0 {
			c.JSON(http.StatusUnprocessableEntity, "detail anggaran belanja tidak boleh kosong")
			c.Abort()
			return
		}

		var sumValue float32
		for i := 0; i < len(project.BudgetDetails); i++ {
			sumValue += project.BudgetDetails[i].Cost
		}
		if sumValue > project.Cost {
			c.JSON(http.StatusUnprocessableEntity, "Jumlah detil biaya tidak boleh melebihi biaya proyek")
			c.Abort()
			return
		}

		c.Next()
	}
}
