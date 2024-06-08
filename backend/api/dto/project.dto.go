package dto

type CreateProject struct {
	Title       string  `json:"title" binding:"required,min=10,max=50"`
	Desc        string  `json:"desc" binding:"required"`
	DeadLine    string  `json:"deadline" binding:"required"`
	Milestone   int8    `json:"milestone" binding:"required,min=1"`
	Tkt         int8    `json:"tkt" binding:"required,min=1,max=12"`
	Cost        float32 `json:"cost" binding:"required,min=0.0"`
	PengajuanID uint    `json:"pengajuan_id" binding:"required"`
}

type ProjectDitolak struct {
	PesanRevisi string `json:"revisi" binding:"required,min=10"`
}

type BudgetDetailsCreate struct {
	Desc    string  `json:"desc" binding:"required"`
	Tahap   uint8   `json:"tahap" binding:"required"`
	Cost    float32 `json:"cost" binding:"required"`
	Percent uint8   `json:"persen" binding:"required"`
	Start   string  `json:"mulai" binding:"required"`
	End     string  `json:"selesai" binding:"required"`
}

type Proposal struct {
	Url string `json:"url" binding:"required"`
}

type Klirens struct {
	Url string `json:"url" binding:"required"`
}