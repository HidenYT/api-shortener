package restapi

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IOutgoingRequestHeaderDAO interface {
	Create(api *OutgoingRequestHeader) error
	Get(id uint) (*OutgoingRequestHeader, error)
	GetAllByConfigID(apiID uint) ([]*OutgoingRequestHeader, error)
	Update(api *OutgoingRequestHeader) error
	Delete(id uint) error
}

type OutgoingRequestHeaderDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestHeaderDAO) Create(api *OutgoingRequestHeader) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Create(api).Error
}

func (dao *OutgoingRequestHeaderDAO) Get(id uint) (*OutgoingRequestHeader, error) {
	result := &OutgoingRequestHeader{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestHeaderDAO) GetAllByConfigID(apiID uint) ([]*OutgoingRequestHeader, error) {
	var result []*OutgoingRequestHeader
	takeResult := dao.db.Where("OutgoingRequestConfigID = ?", apiID).Find(&result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestHeaderDAO) Update(api *OutgoingRequestHeader) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Updates(api).Error
}

func (dao *OutgoingRequestHeaderDAO) Delete(id uint) error {
	return dao.db.Delete(&OutgoingRequestHeader{}, id).Error
}

func NewOutgoingRequestHeaderDAO(conn *gorm.DB, validate *validator.Validate) IOutgoingRequestHeaderDAO {
	return &OutgoingRequestHeaderDAO{db: conn, validate: validate}
}
