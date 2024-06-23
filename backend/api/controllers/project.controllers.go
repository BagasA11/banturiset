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

func (pc *ProjectControllers) OpenDonate(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"pesan": "format page invalid",
		})
		return
	}

	if page <= 0 {
		c.JSON(http.StatusBadRequest, "page tidak boleh 0")
		return
	}

	data, err := pc.Service.OpenDonate(uint(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (pc *ProjectControllers) MyProject(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
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
		"len":  limit,
	})
}

func (pc *ProjectControllers) UploadProposal(c *gin.Context) {

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

	if err := pc.Service.UploadKlirens(uint(id), role_id.(uint), req.Url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "gagal mengunggah proposal",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (pc *ProjectControllers) Preview(c *gin.Context) {
	role_id, _ := c.Get("role_id")

	if role_id.(uint) == 0 {
		c.JSON(http.StatusBadRequest, "id peran diperlukan")
		return
	}

	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "id project invalid",
			"error": err.Error(),
		})
		return
	}

	project, err := pc.Service.Preview(uint(projectID), role_id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})
}

func (pc *ProjectControllers) Reject(c *gin.Context) {

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

func (pc *ProjectControllers) Submit(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "format id pada url invalid")
		return
	}

	role_id, _ := c.Get("role_id")
	if role_id.(uint) == 0 {
		c.JSON(http.StatusBadRequest, "role_id diperlukan")
		return
	}

	if err := pc.Service.SubmitToReviewed(uint(projectID), role_id.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal mensubmit proyek",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (pc *ProjectControllers) Review(c *gin.Context) {

	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "format id tidak benar, harus angka",
			"error": err.Error(),
		})
		return
	}

	project, err := pc.Service.Review(uint(pID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal mengambil data proyek",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"pesan": "sukses",
		"data":  project,
	})
}

func (pc *ProjectControllers) Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "paramater id tidak ditemukan",
			"error": err.Error(),
		})
		return
	}

	p, err := pc.Service.Detail(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "data project tidak ditemukan",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data":  p,
		"pesan": "sukses",
	})
}

func (pc *ProjectControllers) Verfikasi(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id project invalid")
		return
	}

	if _, err := pc.Service.Verifikasi(uint(projectID)); err != nil {
		c.JSON(http.StatusInternalServerError, "gagal memverifikasi proyek")
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (pc *ProjectControllers) Diverifikasi(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "parameter page diperlukan")
		fmt.Println("[controller] error : ", err.Error())
		return
	}

	projects, err := pc.Service.Diverifikasi(uint(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": projects,
		"len":  len(projects),
	})
}

func (pc *ProjectControllers) HasSubmit(c *gin.Context) {
	// validasi di middleware
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "page invalid",
			"error": err.Error(),
		})
		return
	}

	project, err := pc.Service.HasSubmit(uint(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})

}

func (pc *ProjectControllers) Update(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	roleID, _ := c.Get("role_id")
	if roleID == 0 {
		c.JSON(http.StatusBadRequest, "id peneliti diperlukan")
		return
	}

	req := new(dto.EditProject)
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

	if err = pc.Service.Update(uint(projectID), roleID.(uint), *req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, "ok")
}

func (pc *ProjectControllers) OnGoing(c *gin.Context) {
	// tidak perlu hak akses
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "page invalid",
			"error": err.Error(),
		})
		return
	}
	project, err := pc.Service.OnGoing(uint(page))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": project,
	})
}

func (pc *ProjectControllers) Revisi(c *gin.Context) {
	roleID, _ := c.Get("role_id")
	if roleID.(uint) == 0 {
		c.JSON(400, "header id peneliti diperlukan")
		return
	}

	projects, err := pc.Service.Revisi(roleID.(uint))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": projects,
	})
}
