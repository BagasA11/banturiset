package dto

type Pengajuan struct {
	Title       string  `json:"title" binding:"required,min=10"`
	Desc        string  `json:"desc" binding:"required,min=20"`
	LinkWA      string  `json:"linkwa" binding:"required,max=120"`
	LinkPanduan *string `json:"panduan" binding:"max=120"`
	LinkIcon    string  `json:"icon" binding:"required,max=120"`
	Month       uint8   `json:"month" binding:"required,min=1,max=12"`
}
