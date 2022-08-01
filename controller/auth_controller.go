package controller

import (
	"net/http"

	"github.com/ctirouzh/tiny-url/dto"
	"github.com/ctirouzh/tiny-url/pkg/apperr"
	"github.com/ctirouzh/tiny-url/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{
		service: s,
	}
}

func (ctrl *AuthController) SignUp(c *gin.Context) {
	var dto dto.SignUp
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	user, err := ctrl.service.SignUp(&dto)
	if err != nil {
		c.Error(apperr.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var credentials dto.SignIn
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	accessTokenResp, err := ctrl.service.Login(&credentials)
	if err != nil {
		c.Error(apperr.New(http.StatusUnauthorized, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": accessTokenResp})
}
