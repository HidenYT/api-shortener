package api_dao

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OutgoingRequestConfigDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestConfigDAO) Create(config *OutgoingRequestConfig) error {
	err := dao.validate.Struct(config)
	if err != nil {
		return err
	}
	return dao.db.Create(config).Error
}

func (dao *OutgoingRequestConfigDAO) Get(id uint) (*OutgoingRequestConfig, error) {
	result := &OutgoingRequestConfig{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestConfigDAO) GetByAPIID(apiID uint) (*OutgoingRequestConfig, error) {
	result := &OutgoingRequestConfig{}
	takeResult := dao.db.Where("Shortened_API_ID = ?", apiID).Take(&result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestConfigDAO) Update(config *OutgoingRequestConfig) error {
	err := dao.validate.Struct(config)
	if err != nil {
		return err
	}
	return dao.db.Updates(config).Error
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
