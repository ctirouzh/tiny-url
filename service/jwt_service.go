package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/ctirouzh/tiny-url/model"
	"github.com/golang-jwt/jwt"
)

type JwtService struct {
	ttl    time.Duration
	secret string
	issuer string
}

func NewJwtService(ttl time.Duration, secret string, issuer string) *JwtService {
	return &JwtService{
		ttl:    ttl,
		secret: secret,
		issuer: issuer,
	}
}

func (s *JwtService) GenerateJwtToken(user *model.User) (*model.AccessTokenResponse, error) {
	var tokenResp *model.AccessTokenResponse
	claims := model.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(s.ttl).Unix(),
			Issuer:    s.issuer,
		},
		User: &model.UserClaims{
			Username: user.Username,
			UserID:   user.ID.String(),
		},
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenWithClaims.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}
	tokenResp = &model.AccessTokenResponse{
		AccessToken: token,
		TTL:         int(s.ttl.Seconds()),
		ExpiredAt:   time.Now().Add(s.ttl),
		UserID:      user.ID.String(),
	}
	return tokenResp, nil
}

func (s *JwtService) VerifyToken(tokenString string) (*model.AuthClaims, error) {
	claims := &model.AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}
			return []byte(s.secret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}
	claims, ok := token.Claims.(*model.AuthClaims)
	if !ok {
		return nil, errors.New("claims retrieve failed")
	}
	return claims, nil
}
