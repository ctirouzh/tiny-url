package model

import (
	"time"

	"github.com/ctirouzh/tiny-url/dto"
)

type URL struct {
	Hash           string    `json:"hash,omitempty"`
	OriginalURL    string    `json:"original_url,omitempty"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	UserID         string    `json:"user_id"`
}

type URLCache interface {
	SetURL(url *URL)
	GetURL(hash string) *URL
	DeleteURL(hash string) error
}

type URLRepository interface {
	GetUserURLByHash(hash string, user *UserClaims) (*URL, error)
	GetURLByHash(hash string) (*URL, error)
	GetAllURLs(user_id string) ([]URL, error)
	CreateURL(createURLDto *dto.CreateURL, user *User) (*URL, error)
	DeleteURL(hash string, user_id string) error
}
