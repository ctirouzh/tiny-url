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
	Username string `json:"username"`
	UserID   string `json:"user_id"`
}

type AuthClaims struct {
	jwt.StandardClaims
	User *UserClaims `json:"user"`
}
type AuthService interface {
	SignUp() (*User, error)
	Login(credentials *dto.SignIn) (*AccessTokenResponse, error)
}

type JwtService interface {
	GenerateJwtToken(user *User) (*AccessTokenResponse, error)
	VerifyToken(tokenString string) (*AuthClaims, error)
}
