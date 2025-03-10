package restapi

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IOutgoingRequestHeaderDAO interface {
	Create(api *OutgoingRequestHeader) error
	Get(id uint) (*OutgoingRequestHeader, error)
	GetAllByConfigID(configID uint) ([]*OutgoingRequestHeader, error)
	Update(api *OutgoingRequestHeader) error
	Delete(id uint) error
}

type OutgoingRequestHeaderDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *OutgoingRequestHeaderDAO) Create(header *OutgoingRequestHeader) error {
	err := dao.validate.Struct(header)
	if err != nil {
		return err
	}
	return dao.db.Create(header).Error
}

func (dao *OutgoingRequestHeaderDAO) Get(id uint) (*OutgoingRequestHeader, error) {
	result := &OutgoingRequestHeader{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestHeaderDAO) GetAllByConfigID(configID uint) ([]*OutgoingRequestHeader, error) {
	var result []*OutgoingRequestHeader
	takeResult := dao.db.Where("Outgoing_Request_Config_ID = ?", configID).Find(&result)
	return result, takeResult.Error
}

func (dao *OutgoingRequestHeaderDAO) Update(header *OutgoingRequestHeader) error {
	err := dao.validate.Struct(header)
	if err != nil {
		return err
	}
	return dao.db.Updates(header).Error
}

func (dao *OutgoingRequestHeaderDAO) Delete(id uint) error {
	header, err := dao.Get(id)
	if err != nil {
		return err
	}
	return dao.db.Unscoped().Delete(header).Error
}

func NewOutgoingRequestHeaderDAO(conn *gorm.DB, validate *validator.Validate) IOutgoingRequestHeaderDAO {
	return &OutgoingRequestHeaderDAO{db: conn, validate: validate}
}
