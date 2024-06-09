package controllers

import (
	"errors"
	"fmt"
	"mime/multipart"
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

	// validate file size
	if _, err := helpers.ValidateFileSize(c.Request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "ukuran file melebihi kapasitas",
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

func UploadMulti(c *gin.Context) {
	folders := c.QueryArray("folders")
	if len(folders) <= 1 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"pesan":   "sebaiknya lakukan request pada /api/upload?folder=folder",
			"folders": folders,
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	files := form.File["files[]"]

	dest, err := multi_file(folders, files)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"pesan": "gagal melakukan index pada file",
			"error": err.Error(),
		})
		return
	}

	for i, f := range files {
		if err := c.SaveUploadedFile(f, dest[i]); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"pesan": "upload sukses",
		"url":   dest,
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

// /upload?folder=icon,panduan
// /upload?folder=proposal,klirens

func multi_file(folders []string, files []*multipart.FileHeader) ([]string, error) {
	var filenames = []string{}
	if uint(len(folders)) <= 1 {
		return []string{}, errors.New("length is too less")
	}
	if uint(len(files)) <= 1 {
		return []string{}, errors.New("length is too less")
	}
	if uint(len(files)) != uint(len(folders)) {
		return []string{}, errors.New("folders and file must have a exactly same number")
	}

	for i, file := range files {
		f, err := generateFileName(folders[i], file.Filename)
		if err != nil {

			return []string{}, err
		}

		filenames[i] = f
	}

	return filenames, nil
}

func Download(c *gin.Context) {
	println(c.Query("path"))
}
