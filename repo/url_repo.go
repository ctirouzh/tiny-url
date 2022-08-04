package repo

import (
	"context"
	"errors"
	"time"

	"github.com/ctirouzh/tiny-url/dto"
	"github.com/ctirouzh/tiny-url/model"
	"github.com/gocql/gocql"
	"github.com/teris-io/shortid"
)

type URLRespository struct {
	session       *gocql.Session
	cacheRepo     model.URLCache
	defaultUrlTTL time.Duration
}

func NewURLRepository(s *gocql.Session, c model.URLCache, urlTTL time.Duration) *URLRespository {
	return &URLRespository{
		session:       s,
		cacheRepo:     c,
		defaultUrlTTL: urlTTL,
	}
}

func (r *URLRespository) GetURLByHash(hash string) (*model.URL, error) {
	var url *model.URL
	url = r.cacheRepo.GetURL(hash)
	if url != nil {
		return url, nil
	}
	row := map[string]interface{}{}
	var found bool = false
	query := "SELECT * FROM urls WHERE hash = ? LIMIT 1 ALLOW FILTERING"
	iter := r.session.Query(query, hash).Iter()
	defer iter.Close()
	for iter.MapScan(row) {
		found = true
		url = &model.URL{
			Hash:           row["hash"].(string),
			OriginalURL:    row["original_url"].(string),
			CreationDate:   row["creation_date"].(time.Time),
			ExpirationDate: row["expiration_date"].(time.Time),
			UserID:         row["user_id"].(string),
		}
	}
	if !found {
		return nil, errors.New("url not found")
	}
	r.cacheRepo.SetURL(url)
	return url, nil
}

func (r *URLRespository) GetUserURLByHash(hash string, user *model.UserClaims) (*model.URL, error) {
	var url *model.URL
	url = r.cacheRepo.GetURL(hash)
	if url != nil && url.UserID == user.UserID {
		return url, nil
	}
	row := map[string]interface{}{}
	var found bool = false
	query := "SELECT * FROM urls WHERE user_id = ? AND hash = ? LIMIT 1 ALLOW FILTERING"
	iter := r.session.Query(query, user.UserID, hash).Iter()
	defer iter.Close()
	for iter.MapScan(row) {
		found = true
		url = &model.URL{
			Hash:           row["hash"].(string),
			OriginalURL:    row["original_url"].(string),
			CreationDate:   row["creation_date"].(time.Time),
			ExpirationDate: row["expiration_date"].(time.Time),
			UserID:         row["user_id"].(string),
		}
	}
	if !found {
		return nil, errors.New("url not found")
	}
	r.cacheRepo.SetURL(url)
	return url, nil
}

func (r *URLRespository) GetAllURLs(user_id string) ([]model.URL, error) {
	var urls []model.URL

	query := `SELECT * FROM urls WHERE user_id = ?`
	iter := r.session.Query(query, user_id).Iter()
	defer iter.Close()
	for {
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			break
		}
		urls = append(urls, model.URL{
			Hash:           row["hash"].(string),
			OriginalURL:    row["original_url"].(string),
			CreationDate:   row["creation_date"].(time.Time),
			ExpirationDate: row["expiration_date"].(time.Time),
			UserID:         row["user_id"].(string),
		})
	}
	return urls, nil
}

func (r *URLRespository) CreateURL(createURLDto *dto.CreateURL, user *model.User) (*model.URL, error) {

	hash, err := shortid.Generate()
	if err != nil {
		return nil, errors.New("can't generate new hash")
	}

	var count int
	query := `SELECT COUNT(*) FROM urls WHERE user_id = ? AND original_url = ? ALLOW FILTERING`
	r.session.Query(query, user.ID.String(), createURLDto.OriginalURL).Iter().Scan(&count)
	if count > 0 {
		return nil, errors.New("url already hashed")
	}

	tinyurl := &model.URL{
		Hash:           hash,
		OriginalURL:    createURLDto.OriginalURL,
		CreationDate:   time.Now(),
		ExpirationDate: time.Now().Add(r.defaultUrlTTL),
		UserID:         user.ID.String(),
	}
	ctx := context.Background()
	if err := r.session.Query(
		`INSERT INTO urls (hash, original_url, creation_date, expiration_date, user_id) VALUES (?, ?, ?, ?, ?)`,
		tinyurl.Hash, tinyurl.OriginalURL, tinyurl.CreationDate, tinyurl.ExpirationDate, tinyurl.UserID,
	).WithContext(ctx).Exec(); err != nil {
		return nil, err
	}
	return tinyurl, nil
}

func (r *URLRespository) DeleteURL(hash string, user_id string) error {
	query := `DELETE FROM urls WHERE hash = ? AND user_id = ?`
	if err := r.session.Query(query, hash, user_id).Exec(); err != nil {
		return err
	}

	if r.cacheRepo.GetURL(hash) == nil {
		return nil
	}
	if err := r.cacheRepo.DeleteURL(hash); err != nil {
		return err
	}
	return nil
}
