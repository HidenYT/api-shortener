package shortreq

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IShortenedAPIDAO interface {
	Create(api *ShortenedAPI) error
	Get(id uint) (*ShortenedAPI, error)
	Update(api *ShortenedAPI) error
	Delete(id uint) error
}

type ShortenedAPIDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *ShortenedAPIDAO) Create(api *ShortenedAPI) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Create(api).Error
}

func (dao *ShortenedAPIDAO) Get(id uint) (*ShortenedAPI, error) {
	result := &ShortenedAPI{}
	takeResult := dao.db.
		Preload("Config").
		Preload("Config.Headers").
		Preload("Config.Params").
		Preload("ShorteningRules").
		Where("ID = ?", id).
		Take(result)
	return result, takeResult.Error
}

func (dao *ShortenedAPIDAO) Update(api *ShortenedAPI) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(api.Config).Association("Headers").Replace(api.Config.Headers)
		if err != nil {
			return err
		}
		err = tx.Model(api.Config).Association("Params").Replace(api.Config.Params)
		if err != nil {
			return err
		}
		err = tx.Model(api).Association("ShorteningRules").Replace(api.ShorteningRules)
		if err != nil {
			return err
		}
		return tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(api).Error
	})
}

func (dao *ShortenedAPIDAO) Delete(id uint) error {
	return dao.db.Delete(&ShortenedAPI{}, id).Error
}

func NewShortenedAPIDAO(conn *gorm.DB, validate *validator.Validate) IShortenedAPIDAO {
	return &ShortenedAPIDAO{db: conn, validate: validate}
}
