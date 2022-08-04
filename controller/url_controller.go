package controller

import (
	"net/http"

	"github.com/ctirouzh/tiny-url/dto"
	"github.com/ctirouzh/tiny-url/model"
	"github.com/ctirouzh/tiny-url/pkg/apperr"
	"github.com/ctirouzh/tiny-url/service"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type URLController struct {
	service *service.URLService
}

func NewURLController(s *service.URLService) *URLController {
	return &URLController{
		service: s,
	}
}

func (ctrl *URLController) GetAllURLs(c *gin.Context) {
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.Error(apperr.New(http.StatusUnauthorized, "user claims not found"))
		return
	}
	claims := userClaims.(*model.UserClaims)
	urls, err := ctrl.service.GetAllURLs(claims.UserID)
	if err != nil {
		c.Error(apperr.New(http.StatusNotFound, err.Error()))
	}
	c.JSON(http.StatusOK, gin.H{"data": urls})
}

func (ctrl *URLController) CreateURL(c *gin.Context) {

	var createURLDto dto.CreateURL
	if err := c.ShouldBindJSON(&createURLDto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.Error(apperr.New(http.StatusUnauthorized, "user claims not found"))
		return
	}
	user := userClaims.(*model.UserClaims)
	uuid, err := gocql.ParseUUID(user.UserID)
	if err != nil {
		c.Error(apperr.New(http.StatusUnauthorized, "invalid user claims"))
		return
	}
	url, err := ctrl.service.CreateURL(&createURLDto, &model.User{ID: uuid, Username: user.Username})
	if err != nil {
		c.Error(apperr.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": url})
}

func (ctrl *URLController) GetURLByHash(c *gin.Context) {

	var dto dto.GetURLByHash
	if err := c.ShouldBindUri(&dto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.Error(apperr.New(http.StatusUnauthorized, "user claims not found"))
		return
	}
	user := userClaims.(*model.UserClaims)
	url, err := ctrl.service.GetUserURLByHash(dto.Hash, user)
	if err != nil {
		c.Error(apperr.New(http.StatusNotFound, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": url})
}

func (ctrl *URLController) RedirectURLByHash(c *gin.Context) {
	var dto dto.GetURLByHash
	if err := c.ShouldBindUri(&dto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	url, err := ctrl.service.GetURLByHash(dto.Hash)
	if err != nil {
		c.Error(apperr.New(http.StatusNotFound, err.Error()))
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

func (ctrl *URLController) DeleteURL(c *gin.Context) {
	var dto dto.GetURLByHash
	if err := c.ShouldBindUri(&dto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	userClaims, exists := c.Get("user")

	if !exists {
		c.Error(apperr.New(http.StatusUnauthorized, "user claims not found"))
		return
	}
	user := userClaims.(*model.UserClaims)
	uuid, err := gocql.ParseUUID(user.UserID)
	if err != nil {
		c.Error(apperr.New(http.StatusUnauthorized, "invalid user claim"))
		return
	}
	url, err := ctrl.service.GetUserURLByHash(dto.Hash, user)
	if err != nil {
		c.Error(apperr.New(http.StatusNotFound, "url not found"))
		return
	}
	if err := ctrl.service.DeleteURL(url.Hash, uuid.String()); err != nil {
		c.Error(apperr.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
