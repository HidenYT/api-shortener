package restapi

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type IShorteningRuleDAO interface {
	Create(api *ShorteningRule) error
	Get(id uint) (*ShorteningRule, error)
	GetAllByAPIID(apiID uint) ([]*ShorteningRule, error)
	Update(api *ShorteningRule) error
	Delete(id uint) error
}

type ShorteningRuleDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *ShorteningRuleDAO) Create(api *ShorteningRule) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Create(api).Error
}

func (dao *ShorteningRuleDAO) Get(id uint) (*ShorteningRule, error) {
	result := &ShorteningRule{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *ShorteningRuleDAO) GetAllByAPIID(apiID uint) ([]*ShorteningRule, error) {
	var result []*ShorteningRule
	takeResult := dao.db.Where("Shortened_API_ID = ?", apiID).Find(&result)
	return result, takeResult.Error
}

func (dao *ShorteningRuleDAO) Update(api *ShorteningRule) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Updates(api).Error
}

func (dao *ShorteningRuleDAO) Delete(id uint) error {
	rule, err := dao.Get(id)
	if err != nil {
		return err
	}
	return dao.db.Unscoped().Delete(rule).Error
}

func NewShorteningRuleDAO(conn *gorm.DB, validate *validator.Validate) IShorteningRuleDAO {
	return &ShorteningRuleDAO{db: conn, validate: validate}
}
