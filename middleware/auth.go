package middleware

import (
	"net/http"
	"strings"

	"github.com/ctirouzh/tiny-url/pkg/apperr"
	"github.com/ctirouzh/tiny-url/service"
	"github.com/gin-gonic/gin"
)

type authorizeHeader struct {
	Token string `header:"Authorization"`
}

func AuthorizeWithJwt(jwtService *service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authorizeHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.AbortWithError(http.StatusUnauthorized, apperr.New(http.StatusUnauthorized, err.Error()))
			return
		}
		token := strings.Split(h.Token, "Bearer ")
		if len(token) < 2 {
			c.AbortWithError(http.StatusUnauthorized, apperr.New(http.StatusUnauthorized, "must provide authorization header with format `Bearer {token}`"))
			return
		}
		claims, err := jwtService.VerifyToken(token[1])
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, apperr.New(http.StatusUnauthorized, err.Error()))
			return
		}
		c.Set("user", claims.User)
	}
}
