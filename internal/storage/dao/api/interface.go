package api_dao

import db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"

type IShortenedAPIDAO interface {
	Create() (*db_model.ShortenedAPI, error)
	Get(id uint) (*db_model.ShortenedAPI, error)
	Delete(id uint) error
}

type IOutgoingRequestConfigDAO interface {
	Create(api *db_model.OutgoingRequestConfig) error
	Get(id uint) (*db_model.OutgoingRequestConfig, error)
	GetByAPIID(apiID uint) (*db_model.OutgoingRequestConfig, error)
	Update(api *db_model.OutgoingRequestConfig) error
	Delete(id uint) error
}

type IOutgoingRequestHeaderDAO interface {
	Create(api *db_model.OutgoingRequestHeader) error
	Get(id uint) (*db_model.OutgoingRequestHeader, error)
	GetAllByConfigID(configID uint) ([]*db_model.OutgoingRequestHeader, error)
	Update(api *db_model.OutgoingRequestHeader) error
	Delete(id uint) error
}

type IOutgoingRequestParamDAO interface {
	Create(api *db_model.OutgoingRequestParam) error
	Get(id uint) (*db_model.OutgoingRequestParam, error)
	GetAllByConfigID(configID uint) ([]*db_model.OutgoingRequestParam, error)
	Update(api *db_model.OutgoingRequestParam) error
	Delete(id uint) error
}

type IShorteningRuleDAO interface {
	Create(api *db_model.ShorteningRule) error
	Get(id uint) (*db_model.ShorteningRule, error)
	GetAllByAPIID(apiID uint) ([]*db_model.ShorteningRule, error)
	Update(api *db_model.ShorteningRule) error
	Delete(id uint) error
}
