package storage

import (
	"api-shortener/shortreq"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OutgoingRequestConfigDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestConfigDAO) Create(api *shortreq.OutgoingRequestConfig) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Create(api).Error
}

func (dao *OutgoingRequestConfigDAO) Get(id uint) (*shortreq.OutgoingRequestConfig, error) {
	result := &shortreq.OutgoingRequestConfig{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestConfigDAO) GetByAPIID(apiID uint) (*shortreq.OutgoingRequestConfig, error) {
	result := &shortreq.OutgoingRequestConfig{}
	takeResult := dao.db.Where("Shortened_API_ID = ?", apiID).Take(&result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestConfigDAO) Update(api *shortreq.OutgoingRequestConfig) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Updates(api).Error
}

func (dao *OutgoingRequestConfigDAO) Delete(id uint) error {
	config, err := dao.Get(id)
	if err != nil {
		return err
	}
	return dao.db.Unscoped().Delete(config).Error
}

func NewOutgoingRequestConfigDAO(conn *gorm.DB, validate *validator.Validate) *OutgoingRequestConfigDAO {
	return &OutgoingRequestConfigDAO{db: conn, validate: validate}
}
