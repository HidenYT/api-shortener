package http

import api_dao "github.com/HidenYT/api-shortener/internal/storage/dao"

type IAPIService interface {
	Create() (*api_dao.ShortenedAPI, error)
	GetByID(id uint) (*api_dao.ShortenedAPI, error)
	Delete(id uint) error
}

type IRequestConfigService interface {
	Create(api *OutgoingRequestConfigRequest) (*api_dao.OutgoingRequestConfig, error)
	GetByID(id uint) (*api_dao.OutgoingRequestConfig, error)
	GetByAPIID(apiID uint) (*api_dao.OutgoingRequestConfig, error)
	Update(id uint, api *OutgoingRequestConfigRequest) (*api_dao.OutgoingRequestConfig, error)
	Delete(id uint) error
}

type IShorteningRuleService interface {
	Create(api *ShorteningRuleRequest) (*api_dao.ShorteningRule, error)
	GetByID(id uint) (*api_dao.ShorteningRule, error)
	GetAllByAPIID(apiID uint) ([]*api_dao.ShorteningRule, error)
	Update(id uint, api *ShorteningRuleRequest) (*api_dao.ShorteningRule, error)
	Delete(id uint) error
}

type IRequestHeaderService interface {
	Create(api *OutgoingRequestHeaderRequest) (*api_dao.OutgoingRequestHeader, error)
	GetByID(id uint) (*api_dao.OutgoingRequestHeader, error)
	GetAllByConfigID(apiID uint) ([]*api_dao.OutgoingRequestHeader, error)
	Update(id uint, api *OutgoingRequestHeaderRequest) (*api_dao.OutgoingRequestHeader, error)
	Delete(id uint) error
}

type IRequestParamService interface {
	Create(api *OutgoingRequestParamRequest) (*api_dao.OutgoingRequestParam, error)
	GetByID(id uint) (*api_dao.OutgoingRequestParam, error)
	GetAllByConfigID(apiID uint) ([]*api_dao.OutgoingRequestParam, error)
	Update(id uint, api *OutgoingRequestParamRequest) (*api_dao.OutgoingRequestParam, error)
	Delete(id uint) error
}
