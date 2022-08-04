package service

import (
	"github.com/ctirouzh/tiny-url/dto"
	"github.com/ctirouzh/tiny-url/model"
)

type URLService struct {
	urlRepo model.URLRepository
}

func NewUrlService(r model.URLRepository) *URLService {
	return &URLService{
		urlRepo: r,
	}
}

func (s *URLService) GetAllURLs(user_id string) ([]model.URL, error) {
	return s.urlRepo.GetAllURLs(user_id)
}

func (s *URLService) CreateURL(createURLDto *dto.CreateURL, user *model.User) (*model.URL, error) {
	return s.urlRepo.CreateURL(createURLDto, user)
}

func (s *URLService) GetUserURLByHash(hash string, user *model.UserClaims) (*model.URL, error) {
	return s.urlRepo.GetUserURLByHash(hash, user)
}

func (s *URLService) GetURLByHash(hash string) (*model.URL, error) {
	return s.urlRepo.GetURLByHash(hash)
}

func (s *URLService) DeleteURL(hash string, user_id string) error {
	return s.urlRepo.DeleteURL(hash, user_id)
}
