package storage

import (
	"api-shortener/shortreq"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ShortenedAPIDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *ShortenedAPIDAO) Create() (*shortreq.ShortenedAPI, error) {
	api := &shortreq.ShortenedAPI{}
	return api, dao.db.Create(api).Error
}

func (dao *ShortenedAPIDAO) Get(id uint) (*shortreq.ShortenedAPI, error) {
	result := &shortreq.ShortenedAPI{}
	takeResult := dao.db.Where("ID = ?", id).Preload("ShorteningRules").Take(result)
	return result, takeResult.Error
}

func (dao *ShortenedAPIDAO) Update(api *shortreq.ShortenedAPI) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Updates(api).Error
}

func (dao *ShortenedAPIDAO) Delete(id uint) error {
	api, err := dao.Get(id)
	if err != nil {
		return err
	}
	return dao.db.Unscoped().Delete(&api).Error
}

func NewShortenedAPIDAO(conn *gorm.DB, validate *validator.Validate) *ShortenedAPIDAO {
	return &ShortenedAPIDAO{db: conn, validate: validate}
}
