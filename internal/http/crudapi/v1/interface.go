package http

import db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"

type IAPIService interface {
	Create() (*db_model.ShortenedAPI, error)
	GetByID(id uint) (*db_model.ShortenedAPI, error)
	Delete(id uint) error
}

type IRequestConfigService interface {
	Create(api *OutgoingRequestConfigRequest) (*db_model.OutgoingRequestConfig, error)
	GetByID(id uint) (*db_model.OutgoingRequestConfig, error)
	GetByAPIID(apiID uint) (*db_model.OutgoingRequestConfig, error)
	Update(id uint, api *OutgoingRequestConfigRequest) (*db_model.OutgoingRequestConfig, error)
	Delete(id uint) error
}

type IShorteningRuleService interface {
	Create(api *ShorteningRuleRequest) (*db_model.ShorteningRule, error)
	GetByID(id uint) (*db_model.ShorteningRule, error)
	GetAllByAPIID(apiID uint) ([]*db_model.ShorteningRule, error)
	Update(id uint, api *ShorteningRuleRequest) (*db_model.ShorteningRule, error)
	Delete(id uint) error
}

type IRequestHeaderService interface {
	Create(api *OutgoingRequestHeaderRequest) (*db_model.OutgoingRequestHeader, error)
	GetByID(id uint) (*db_model.OutgoingRequestHeader, error)
	GetAllByConfigID(apiID uint) ([]*db_model.OutgoingRequestHeader, error)
	Update(id uint, api *OutgoingRequestHeaderRequest) (*db_model.OutgoingRequestHeader, error)
	Delete(id uint) error
}

type IRequestParamService interface {
	Create(api *OutgoingRequestParamRequest) (*db_model.OutgoingRequestParam, error)
	GetByID(id uint) (*db_model.OutgoingRequestParam, error)
	GetAllByConfigID(apiID uint) ([]*db_model.OutgoingRequestParam, error)
	Update(id uint, api *OutgoingRequestParamRequest) (*db_model.OutgoingRequestParam, error)
	Delete(id uint) error
}
