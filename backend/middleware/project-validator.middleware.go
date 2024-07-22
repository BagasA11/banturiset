package middleware

import (
	"net/http"

	"strconv"

	"github.com/bagasa11/banturiset/api/services"
	"github.com/gin-gonic/gin"
)

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

		if project.ProposalUrl == nil || project.KlirensUrl == nil {
			c.JSON(http.StatusUnprocessableEntity, "proposal dan surat klirens penelitian tidak boleh kosong")
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

		c.Next()
	}
}
