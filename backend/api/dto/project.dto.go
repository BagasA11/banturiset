package dto

import "time"

type CreateProject struct {
	Title       string  `json:"title" binding:"required,min=10,max=50"`
	Desc        string  `json:"desc" binding:"required"`
	Year        int16   `json:"year" binding:"required"`
	Milestone   int8    `json:"milestone" binding:"required,min=1"`
	Tkt         int8    `json:"tkt" binding:"required,min=1,max=12"`
	Cost        float32 `json:"cost" binding:"required,min=0.0"`
	PengajuanID uint    `json:"pengajuan_id" binding:"required"`
}

type EditProject struct {
	Desc      string  `json:"desc" binding:"required"`
	Milestone int8    `json:"milestone" binding:"required,min=1"`
	Tkt       int8    `json:"tkt" binding:"required,min=1,max=12"`
	Cost      float32 `json:"cost" binding:"required,min=0.0"`
}

type ProjectDitolak struct {
	PesanRevisi string `json:"revisi" binding:"required,min=10"`
}

type BudgetDetailsCreate struct {
	Desc string `json:"desc" binding:"required"`

	Cost float32 `json:"cost" binding:"required"`
}

type Proposal struct {
	Url string `json:"url" binding:"required"`
}

type Klirens struct {
	Url string `json:"url" binding:"required"`
}

type TahapCreate struct {
	CostPercent uint8  `json:"percent" binding:"required"`
	Tahap       uint8  `json:"tahap" binding:"required"`
	Start       string `json:"awal" binding:"required"`
	End         string `json:"akhir" binding:"required"`
}

type CreateDonasi struct {
	Jml    float32 `json:"jml"`
	Method string  `json:"method"`
}

type InvoicePage struct {
	ID        string
	Jml       float32
	Fee       float32
	CreatedAt time.Time
}

type NotifInvoice struct {
	ExternalID string  `json:"external_id"`
	Status     string  `json:"status"`
	Amount     float32 `json:"amount"`
}
