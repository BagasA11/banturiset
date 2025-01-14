package dto

import "mime/multipart"

type UserRegister struct {
	FName    string `json:"fname" binding:"required,min=10"`
	Email    string `json:"email" binding:"required,min=8" `
	Role     string `json:"role" binding:"required" `
	Password string `json:"password" binding:"required,min=8"`
	PostCode string `json:"post" binding:"required,min=4"`
}

type Profile struct {
	File multipart.FileHeader `form:"file" binding:"required"`
}

type PenelitiRegister struct {
	NIP       string `json:"nip" binding:"required,min=5" `
	Role      string `json:"role" binding:"required" `
	Institute string `json:"institute" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type DonaturRegister struct {
	Role string `json:"role" binding:"required"`
}

type ReviewRegister struct {
	NIP string `json:"nip" binding:"required,min=5"`
}

type Login struct {
	Email    string `json:"email" binding:"required,min=8" `
	Password string `json:"password" binding:"required,min=8"`
}
