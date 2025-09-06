package api_dao

import (
	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OutgoingRequestConfigDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestConfigDAO) Create(config *db_model.OutgoingRequestConfig) error {
	err := dao.validate.Struct(config)
	if err != nil {
		return err
	}
	return dao.db.Create(config).Error
}

func (dao *OutgoingRequestConfigDAO) Get(id uint) (*db_model.OutgoingRequestConfig, error) {
	result := &db_model.OutgoingRequestConfig{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestConfigDAO) GetByAPIID(apiID uint) (*db_model.OutgoingRequestConfig, error) {
	result := &db_model.OutgoingRequestConfig{}
	takeResult := dao.db.Where("Shortened_API_ID = ?", apiID).Take(&result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestConfigDAO) Update(config *db_model.OutgoingRequestConfig) error {
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
