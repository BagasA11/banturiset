package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/bagasa11/banturiset/helpers"
	"github.com/gin-gonic/gin"
)

const pdf string = "pdf"
const jpg = "jpg"
const jpeg = "jpeg"
const png = "png"

func Upload(c *gin.Context) {

	folder := c.Query("folder")
	if folder == "" {
		c.JSON(http.StatusBadRequest, "tambahkan parameter folder")
		return
	}

	file_header, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pesan": "gagal mengambil file dari form html",
			"error": err.Error(),
		})
		return
	}

	newname, err := generateFileName(folder, file_header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := c.SaveUploadedFile(file_header, newname); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
			"pesan": "gagal mengupload file",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pesan": "upload file sukses",
		"url":   newname,
	})
}

func generateFileName(folder string, filename string) (string, error) {

	foldername := strings.ToLower(folder)

	if len(strings.Split(filename, ".")) != 2 {
		return "", errors.New("format extensi file invalid. contoh valid: proposal.pdf")
	}

	ext := strings.ToLower(strings.Split(filename, ".")[1])

	if !slices.Contains([]string{pdf, jpeg, jpg, png}, strings.ToLower(ext)) {
		return "", errors.New("format file bukan jpg, jpeg, png, atau pdf")
	}

	if slices.Contains([]string{"proposal", "klirens", "laporan", "panduan"}, foldername) && ext != pdf {
		return "", errors.New("format file harus pdf")
	}
	if slices.Contains([]string{"gambar", "profil", "icon"}, foldername) && !slices.Contains([]string{png, jpeg, jpg}, ext) {
		return "", errors.New("format file harus png, jpg, atau jpeg")
	}
	// dir/proposal/xefe1nfem.pdf
	return fmt.Sprintf("./file/%s/%s.%s", foldername, helpers.RandStr(7), ext), nil
}
