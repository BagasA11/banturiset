package dto

type ProjectCreate struct {
	Title     string  `json:"title" binding:"required,min=10,max=50"`
	Desc      string  `json:"desc" binding:"required"`
	DeadLine  string  `json:"deadline" binding:"required"`
	Milestone int8    `json:"milestone" binding:"required,min=1"`
	Tkt       int8    `json:"tkt" binding:"required,min=1"`
	Cost      float32 `json:"cost" binding:"required,min=0.0"`
}

type ProjectDitolak struct {
	PesanRevisi string `json:"revisi" binding:"required,min=10"`
}
