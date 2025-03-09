package restapi

import (
	"gorm.io/gorm"
)

type IRESTService interface {
	CreateAPI(api *ShortenedAPI) (*ShortenedAPI, error)
	DeleteAPI(id uint) error

	CreateConfig(api *OutgoingRequestConfig) (*OutgoingRequestConfig, error)
	GetConfig(id uint) (*OutgoingRequestConfig, error)
	GetAllConfigsByAPIID(apiID uint) ([]*OutgoingRequestConfig, error)
	UpdateConfig(api *OutgoingRequestConfig) (*OutgoingRequestConfig, error)
	DeleteConfig(id uint) error

	CreateRule(api *ShorteningRule) (*ShorteningRule, error)
	GetRule(id uint) (*ShorteningRule, error)
	GetAllRulesByAPIID(apiID uint) ([]*ShorteningRule, error)
	UpdateRule(api *ShorteningRule) (*ShorteningRule, error)
	DeleteRule(id uint) error

	CreateRequestHeader(api *OutgoingRequestHeader) (*OutgoingRequestHeader, error)
	GetRequestHeader(id uint) (*OutgoingRequestHeader, error)
	GetAllRequestHeadersByConfigID(apiID uint) ([]*OutgoingRequestHeader, error)
	UpdateRequestHeader(api *OutgoingRequestHeader) (*OutgoingRequestHeader, error)
	DeleteRequestHeader(id uint) error

	CreateRequestParam(api *OutgoingRequestParam) (*OutgoingRequestParam, error)
	GetRequestParam(id uint) (*OutgoingRequestParam, error)
	GetAllRequestParamsByConfigID(apiID uint) ([]*OutgoingRequestParam, error)
	UpdateRequestParam(api *OutgoingRequestParam) (*OutgoingRequestParam, error)
	DeleteRequestParam(id uint) error
}

type RESTService struct {
	db               *gorm.DB
	apiDAO           IShortenedAPIDAO
	requestConfigDAO IOutgoingRequestConfigDAO
	requestHeaderDAO IOutgoingRequestHeaderDAO
	requestParamDAO  IOutgoingRequestParamDAO
	ruleDAO          IShorteningRuleDAO
}

func NewRESTService(
	db *gorm.DB,
	apiDAO IShortenedAPIDAO,
	requestConfigDAO IOutgoingRequestConfigDAO,
	requestHeaderDAO IOutgoingRequestHeaderDAO,
	requestParamDAO IOutgoingRequestParamDAO,
	ruleDAO IShorteningRuleDAO,
) IRESTService {
	return &RESTService{
		db:               db,
		apiDAO:           apiDAO,
		requestConfigDAO: requestConfigDAO,
		requestHeaderDAO: requestHeaderDAO,
		requestParamDAO:  requestParamDAO,
		ruleDAO:          ruleDAO,
	}
}

// API

func (s *RESTService) CreateAPI(api *ShortenedAPI) (*ShortenedAPI, error) {
	api.ID = 0
	err := s.apiDAO.Create(api)
	return api, err
}

func (s *RESTService) DeleteAPI(id uint) error {
	return s.apiDAO.Delete(id)
}

// RequestConfig

func (s *RESTService) CreateConfig(config *OutgoingRequestConfig) (*OutgoingRequestConfig, error) {
	config.ID = 0
	err := s.requestConfigDAO.Create(config)
	return config, err
}

func (s *RESTService) GetConfig(id uint) (*OutgoingRequestConfig, error) {
	return s.requestConfigDAO.Get(id)
}

func (s *RESTService) GetAllConfigsByAPIID(id uint) ([]*OutgoingRequestConfig, error) {
	return s.requestConfigDAO.GetAllByAPIID(id)
}

func (s *RESTService) UpdateConfig(config *OutgoingRequestConfig) (*OutgoingRequestConfig, error) {
	err := s.requestConfigDAO.Update(config)
	return config, err
}

func (s *RESTService) DeleteConfig(id uint) error {
	return s.requestConfigDAO.Delete(id)
}

// ShorteningRule

func (s *RESTService) CreateRule(rule *ShorteningRule) (*ShorteningRule, error) {
	rule.ID = 0
	err := s.ruleDAO.Create(rule)
	return rule, err
}

func (s *RESTService) GetRule(id uint) (*ShorteningRule, error) {
	return s.ruleDAO.Get(id)
}

func (s *RESTService) GetAllRulesByAPIID(id uint) ([]*ShorteningRule, error) {
	return s.ruleDAO.GetAllByAPIID(id)
}

func (s *RESTService) UpdateRule(rule *ShorteningRule) (*ShorteningRule, error) {
	err := s.ruleDAO.Update(rule)
	return rule, err
}

func (s *RESTService) DeleteRule(id uint) error {
	return s.ruleDAO.Delete(id)
}

// Header

func (s *RESTService) CreateRequestHeader(header *OutgoingRequestHeader) (*OutgoingRequestHeader, error) {
	header.ID = 0
	err := s.requestHeaderDAO.Create(header)
	return header, err
}

func (s *RESTService) GetRequestHeader(id uint) (*OutgoingRequestHeader, error) {
	return s.requestHeaderDAO.Get(id)
}

func (s *RESTService) GetAllRequestHeadersByConfigID(id uint) ([]*OutgoingRequestHeader, error) {
	return s.requestHeaderDAO.GetAllByConfigID(id)
}

func (s *RESTService) UpdateRequestHeader(header *OutgoingRequestHeader) (*OutgoingRequestHeader, error) {
	err := s.requestHeaderDAO.Update(header)
	return header, err
}

func (s *RESTService) DeleteRequestHeader(id uint) error {
	return s.requestHeaderDAO.Delete(id)
}

// Param

func (s *RESTService) CreateRequestParam(param *OutgoingRequestParam) (*OutgoingRequestParam, error) {
	param.ID = 0
	err := s.requestParamDAO.Create(param)
	return param, err
}

func (s *RESTService) GetRequestParam(id uint) (*OutgoingRequestParam, error) {
	return s.requestParamDAO.Get(id)
}

func (s *RESTService) GetAllRequestParamsByConfigID(id uint) ([]*OutgoingRequestParam, error) {
	return s.requestParamDAO.GetAllByConfigID(id)
}

func (s *RESTService) UpdateRequestParam(param *OutgoingRequestParam) (*OutgoingRequestParam, error) {
	err := s.requestParamDAO.Update(param)
	return param, err
}

func (s *RESTService) DeleteRequestParam(id uint) error {
	return s.requestParamDAO.Delete(id)
}
