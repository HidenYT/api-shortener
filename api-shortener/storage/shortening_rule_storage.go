package storage

import (
	"api-shortener/shortreq"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ShorteningRuleDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *ShorteningRuleDAO) Create(api *shortreq.ShorteningRule) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Create(api).Error
}

func (dao *ShorteningRuleDAO) Get(id uint) (*shortreq.ShorteningRule, error) {
	result := &shortreq.ShorteningRule{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *ShorteningRuleDAO) GetAllByAPIID(apiID uint) ([]*shortreq.ShorteningRule, error) {
	var result []*shortreq.ShorteningRule
	takeResult := dao.db.Where("Shortened_API_ID = ?", apiID).Find(&result)
	return result, takeResult.Error
}

func (dao *ShorteningRuleDAO) Update(api *shortreq.ShorteningRule) error {
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

func NewShorteningRuleDAO(conn *gorm.DB, validate *validator.Validate) *ShorteningRuleDAO {
	return &ShorteningRuleDAO{db: conn, validate: validate}
}
