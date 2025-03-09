package restapi

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IOutgoingRequestParamDAO interface {
	Create(api *OutgoingRequestParam) error
	Get(id uint) (*OutgoingRequestParam, error)
	GetAllByConfigID(apiID uint) ([]*OutgoingRequestParam, error)
	Update(api *OutgoingRequestParam) error
	Delete(id uint) error
}

type OutgoingRequestParamDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestParamDAO) Create(api *OutgoingRequestParam) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Create(api).Error
}

func (dao *OutgoingRequestParamDAO) Get(id uint) (*OutgoingRequestParam, error) {
	result := &OutgoingRequestParam{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestParamDAO) GetAllByConfigID(apiID uint) ([]*OutgoingRequestParam, error) {
	var result []*OutgoingRequestParam
	takeResult := dao.db.Where("OutgoingRequestConfigID = ?", apiID).Find(&result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestParamDAO) Update(api *OutgoingRequestParam) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Updates(api).Error
}

func (dao *OutgoingRequestParamDAO) Delete(id uint) error {
	return dao.db.Delete(&OutgoingRequestParam{}, id).Error
}

func NewOutgoingRequestParamDAO(conn *gorm.DB, validate *validator.Validate) IOutgoingRequestParamDAO {
	return &OutgoingRequestParamDAO{db: conn, validate: validate}
}
