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

func (ctrl *URLController) CreateURL(c *gin.Context) {

	var createURLDto dto.CreateURL
	if err := c.ShouldBindJSON(&createURLDto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.Error(apperr.New(http.StatusBadRequest, "User Claims Not Found"))
		return
	}
	user := userClaims.(*model.UserClaims)
	uuid, err := gocql.ParseUUID(user.UserID)
	if err != nil {
		c.Error(apperr.New(http.StatusBadRequest, "User Claims Invalid"))
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

	var getURLByHashDto dto.GetURLByHash
	if err := c.ShouldBindUri(&getURLByHashDto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.Error(apperr.New(http.StatusBadRequest, "User Claims Not Found"))
		return
	}
	user := userClaims.(*model.UserClaims)
	url, err := ctrl.service.GetUserURLByHash(getURLByHashDto.Hash, user)
	if err != nil {
		c.Error(apperr.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": url})
}

func (ctrl *URLController) RedirectURLByHash(c *gin.Context) {
	var getURLByHashDto dto.GetURLByHash
	if err := c.ShouldBindUri(&getURLByHashDto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	url, err := ctrl.service.GetURLByHash(getURLByHashDto.Hash)
	if err != nil {
		c.Error(apperr.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

func (ctrl *URLController) DeleteURL(c *gin.Context) {
	var getURLByHashDto dto.GetURLByHash
	if err := c.ShouldBindUri(&getURLByHashDto); err != nil {
		c.Error(apperr.New(http.StatusBadRequest, err.Error()))
		return
	}
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.Error(apperr.New(http.StatusBadRequest, "User Claims Not Found"))
		return
	}
	user := userClaims.(*model.UserClaims)
	uuid, err := gocql.ParseUUID(user.UserID)
	if err != nil {
		c.Error(apperr.New(http.StatusBadRequest, "User Claims Invalid"))
		return
	}
	url, err := ctrl.service.GetUserURLByHash(getURLByHashDto.Hash, user)
	if err != nil {
		c.Error(apperr.New(http.StatusBadRequest, "url not found"))
		return
	}
	if err := ctrl.service.DeleteURL(url.Hash, uuid.String()); err != nil {
		c.Error(apperr.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
