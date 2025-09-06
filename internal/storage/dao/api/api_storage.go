package api_dao

import (
	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ShortenedAPIDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *ShortenedAPIDAO) Create() (*db_model.ShortenedAPI, error) {
	api := &db_model.ShortenedAPI{}
	return api, dao.db.Create(api).Error
}

func (dao *ShortenedAPIDAO) Get(id uint) (*db_model.ShortenedAPI, error) {
	result := &db_model.ShortenedAPI{}
	takeResult := dao.db.Where("ID = ?", id).Preload("ShorteningRules").Take(result)
	return result, takeResult.Error
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
