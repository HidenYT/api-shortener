package restapi

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IShortenedAPIDAO interface {
	Create() (*ShortenedAPI, error)
	Get(id uint) (*ShortenedAPI, error)
	Update(api *ShortenedAPI) error
	Delete(id uint) error
}

type ShortenedAPIDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *ShortenedAPIDAO) Create() (*ShortenedAPI, error) {
	api := &ShortenedAPI{}
	return api, dao.db.Create(api).Error
}

func (dao *ShortenedAPIDAO) Get(id uint) (*ShortenedAPI, error) {
	result := &ShortenedAPI{}
	takeResult := dao.db.Where("ID = ?", id).Preload("ShorteningRules").Take(result)
	return result, takeResult.Error
}

func (dao *ShortenedAPIDAO) Update(api *ShortenedAPI) error {
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

func NewShortenedAPIDAO(conn *gorm.DB, validate *validator.Validate) IShortenedAPIDAO {
	return &ShortenedAPIDAO{db: conn, validate: validate}
}
