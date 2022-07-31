package dto

type CreateURL struct {
	OriginalURL string `json:"original_url" binding:"required,url"`
}

type GetURLByHash struct {
	Hash string `uri:"hash" binding:"required"`
}
