package restapi

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IOutgoingRequestConfigDAO interface {
	Create(api *OutgoingRequestConfig) error
	Get(id uint) (*OutgoingRequestConfig, error)
	GetByAPIID(apiID uint) (*OutgoingRequestConfig, error)
	Update(api *OutgoingRequestConfig) error
	Delete(id uint) error
}

type OutgoingRequestConfigDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestConfigDAO) Create(api *OutgoingRequestConfig) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Create(api).Error
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

func (dao *OutgoingRequestConfigDAO) Update(api *OutgoingRequestConfig) error {
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

func NewOutgoingRequestConfigDAO(conn *gorm.DB, validate *validator.Validate) IOutgoingRequestConfigDAO {
	return &OutgoingRequestConfigDAO{db: conn, validate: validate}
}
