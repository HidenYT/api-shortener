package api_dao

import (
	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OutgoingRequestParamDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestParamDAO) Create(param *db_model.OutgoingRequestParam) error {
	err := dao.validate.Struct(param)
	if err != nil {
		return err
	}
	return dao.db.Create(param).Error
}

func (dao *OutgoingRequestParamDAO) Get(id uint) (*db_model.OutgoingRequestParam, error) {
	result := &db_model.OutgoingRequestParam{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestParamDAO) GetAllByConfigID(configID uint) ([]*db_model.OutgoingRequestParam, error) {
	var result []*db_model.OutgoingRequestParam
	takeResult := dao.db.Where("Outgoing_Request_Config_ID = ?", configID).Find(&result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestParamDAO) Update(param *db_model.OutgoingRequestParam) error {
	err := dao.validate.Struct(param)
	if err != nil {
		return err
	}
	return dao.db.Updates(param).Error
}

func (dao *OutgoingRequestParamDAO) Delete(id uint) error {
	param, err := dao.Get(id)
	if err != nil {
		return err
	}
	return dao.db.Unscoped().Delete(param).Error
}

func NewOutgoingRequestParamDAO(conn *gorm.DB, validate *validator.Validate) *OutgoingRequestParamDAO {
	return &OutgoingRequestParamDAO{db: conn, validate: validate}
}
