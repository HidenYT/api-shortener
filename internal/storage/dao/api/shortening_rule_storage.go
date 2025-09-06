package api_dao

import (
	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ShorteningRuleDAO struct {
	db       *gorm.DB
	validate *validator.Validate
}

func (dao *ShorteningRuleDAO) Create(api *db_model.ShorteningRule) error {
	err := dao.validate.Struct(api)
	if err != nil {
		return err
	}
	return dao.db.Create(api).Error
}

func (dao *ShorteningRuleDAO) Get(id uint) (*db_model.ShorteningRule, error) {
	result := &db_model.ShorteningRule{}
	takeResult := dao.db.Where("ID = ?", id).Take(result)
	return result, takeResult.Error
}

func (dao *ShorteningRuleDAO) GetAllByAPIID(apiID uint) ([]*db_model.ShorteningRule, error) {
	var result []*db_model.ShorteningRule
	takeResult := dao.db.Where("Shortened_API_ID = ?", apiID).Find(&result)
	return result, takeResult.Error
}

func (dao *ShorteningRuleDAO) Update(api *db_model.ShorteningRule) error {
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
