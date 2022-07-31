package model

import (
	"time"

	"github.com/ctirouzh/tiny-url/dto"
	"github.com/golang-jwt/jwt"
)

type AccessTokenResponse struct {
	AccessToken string    `json:"access_token"`
	TTL         int       `json:"ttl"`
	ExpiredAt   time.Time `json:"expired_at"`
	UserID      string    `json:"user_id"`
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserID   string `json:"user_id"`
}

type AuthService interface {
	SignUp() (*User, error)
	Login(credentials *dto.SignInDto) (*AccessTokenResponse, error)
}

type JwtService interface {
	GenerateJwtToken(user *User) (*AccessTokenResponse, error)
	VerifyToken(tokenString string) (*UserClaims, error)
}
