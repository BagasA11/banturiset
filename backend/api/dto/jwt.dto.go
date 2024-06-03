package dto

type RefreshTokenRequest struct {
	OldToken string `json:"oldToken"`
}
