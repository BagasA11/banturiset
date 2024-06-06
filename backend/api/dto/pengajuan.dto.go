package dto

type Pengajuan struct {
	Title       string  `json:"title" binding:"required,min=10"`
	Desc        string  `json:"desc" binding:"required,min=20"`
	LinkWA      string  `json:"linkwa" binding:"required,max=120"`
	LinkPanduan *string `json:"panduan" binding:"max=120"`
	LinkIcon    string  `json:"icon" binding:"required,max=120"`
	ClosedAt    string  `json:"closed_at" binding:"required"`
}
