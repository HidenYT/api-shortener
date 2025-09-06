package http

import (
	"github.com/HidenYT/api-shortener/shortreq"
)

type IAPIService interface {
	Create() (*shortreq.ShortenedAPI, error)
	GetByID(id uint) (*shortreq.ShortenedAPI, error)
	Delete(id uint) error
}

type IRequestConfigService interface {
	Create(api *OutgoingRequestConfigRequest) (*shortreq.OutgoingRequestConfig, error)
	GetByID(id uint) (*shortreq.OutgoingRequestConfig, error)
	GetByAPIID(apiID uint) (*shortreq.OutgoingRequestConfig, error)
	Update(id uint, api *OutgoingRequestConfigRequest) (*shortreq.OutgoingRequestConfig, error)
	Delete(id uint) error
}

type IShorteningRuleService interface {
	Create(api *ShorteningRuleRequest) (*shortreq.ShorteningRule, error)
	GetByID(id uint) (*shortreq.ShorteningRule, error)
	GetAllByAPIID(apiID uint) ([]*shortreq.ShorteningRule, error)
	Update(id uint, api *ShorteningRuleRequest) (*shortreq.ShorteningRule, error)
	Delete(id uint) error
}

type IRequestHeaderService interface {
	Create(api *OutgoingRequestHeaderRequest) (*shortreq.OutgoingRequestHeader, error)
	GetByID(id uint) (*shortreq.OutgoingRequestHeader, error)
	GetAllByConfigID(apiID uint) ([]*shortreq.OutgoingRequestHeader, error)
	Update(id uint, api *OutgoingRequestHeaderRequest) (*shortreq.OutgoingRequestHeader, error)
	Delete(id uint) error
}

type IRequestParamService interface {
	Create(api *OutgoingRequestParamRequest) (*shortreq.OutgoingRequestParam, error)
	GetByID(id uint) (*shortreq.OutgoingRequestParam, error)
	GetAllByConfigID(apiID uint) ([]*shortreq.OutgoingRequestParam, error)
	Update(id uint, api *OutgoingRequestParamRequest) (*shortreq.OutgoingRequestParam, error)
	Delete(id uint) error
}
