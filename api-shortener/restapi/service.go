package restapi

import (
	"gorm.io/gorm"
)

type IRESTService interface {
	CreateAPI() (*ShortenedAPI, error)
	DeleteAPI(id uint) error

	CreateConfig(api *OutgoingRequestConfigRequest) (*OutgoingRequestConfig, error)
	GetConfig(id uint) (*OutgoingRequestConfig, error)
	GetConfigByAPIID(apiID uint) (*OutgoingRequestConfig, error)
	UpdateConfig(id uint, api *OutgoingRequestConfigRequest) (*OutgoingRequestConfig, error)
	DeleteConfig(id uint) error

	CreateRule(api *ShorteningRuleRequest) (*ShorteningRule, error)
	GetRule(id uint) (*ShorteningRule, error)
	GetAllRulesByAPIID(apiID uint) ([]*ShorteningRule, error)
	UpdateRule(id uint, api *ShorteningRuleRequest) (*ShorteningRule, error)
	DeleteRule(id uint) error

	CreateRequestHeader(api *OutgoingRequestHeaderRequest) (*OutgoingRequestHeader, error)
	GetRequestHeader(id uint) (*OutgoingRequestHeader, error)
	GetAllRequestHeadersByConfigID(apiID uint) ([]*OutgoingRequestHeader, error)
	UpdateRequestHeader(id uint, api *OutgoingRequestHeaderRequest) (*OutgoingRequestHeader, error)
	DeleteRequestHeader(id uint) error

	CreateRequestParam(api *OutgoingRequestParamRequest) (*OutgoingRequestParam, error)
	GetRequestParam(id uint) (*OutgoingRequestParam, error)
	GetAllRequestParamsByConfigID(apiID uint) ([]*OutgoingRequestParam, error)
	UpdateRequestParam(id uint, api *OutgoingRequestParamRequest) (*OutgoingRequestParam, error)
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

func (s *RESTService) CreateAPI() (*ShortenedAPI, error) {
	return s.apiDAO.Create()
}

func (s *RESTService) DeleteAPI(id uint) error {
	return s.apiDAO.Delete(id)
}

// RequestConfig

func (s *RESTService) CreateConfig(configRequest *OutgoingRequestConfigRequest) (*OutgoingRequestConfig, error) {
	config := outgoingRequestConfigRequestToDBModel(configRequest)
	err := s.requestConfigDAO.Create(config)
	return config, err
}

func (s *RESTService) GetConfig(id uint) (*OutgoingRequestConfig, error) {
	return s.requestConfigDAO.Get(id)
}

func (s *RESTService) GetConfigByAPIID(id uint) (*OutgoingRequestConfig, error) {
	return s.requestConfigDAO.GetByAPIID(id)
}

func (s *RESTService) UpdateConfig(id uint, configRequest *OutgoingRequestConfigRequest) (*OutgoingRequestConfig, error) {
	config := outgoingRequestConfigRequestToDBModel(configRequest)
	config.ID = id
	err := s.requestConfigDAO.Update(config)
	return config, err
}

func (s *RESTService) DeleteConfig(id uint) error {
	return s.requestConfigDAO.Delete(id)
}

// ShorteningRule

func (s *RESTService) CreateRule(ruleRequest *ShorteningRuleRequest) (*ShorteningRule, error) {
	rule := shorteningRuleRequestToDBModel(ruleRequest)
	err := s.ruleDAO.Create(rule)
	return rule, err
}

func (s *RESTService) GetRule(id uint) (*ShorteningRule, error) {
	return s.ruleDAO.Get(id)
}

func (s *RESTService) GetAllRulesByAPIID(id uint) ([]*ShorteningRule, error) {
	return s.ruleDAO.GetAllByAPIID(id)
}

func (s *RESTService) UpdateRule(id uint, ruleRequest *ShorteningRuleRequest) (*ShorteningRule, error) {
	rule := shorteningRuleRequestToDBModel(ruleRequest)
	rule.ID = id
	err := s.ruleDAO.Update(rule)
	return rule, err
}

func (s *RESTService) DeleteRule(id uint) error {
	return s.ruleDAO.Delete(id)
}

// Header

func (s *RESTService) CreateRequestHeader(headerRequest *OutgoingRequestHeaderRequest) (*OutgoingRequestHeader, error) {
	header := outgoingRequestHeaderRequestToDBModel(headerRequest)
	err := s.requestHeaderDAO.Create(header)
	return header, err
}

func (s *RESTService) GetRequestHeader(id uint) (*OutgoingRequestHeader, error) {
	return s.requestHeaderDAO.Get(id)
}

func (s *RESTService) GetAllRequestHeadersByConfigID(id uint) ([]*OutgoingRequestHeader, error) {
	return s.requestHeaderDAO.GetAllByConfigID(id)
}

func (s *RESTService) UpdateRequestHeader(id uint, headerRequest *OutgoingRequestHeaderRequest) (*OutgoingRequestHeader, error) {
	header := outgoingRequestHeaderRequestToDBModel(headerRequest)
	header.ID = id
	err := s.requestHeaderDAO.Update(header)
	return header, err
}

func (s *RESTService) DeleteRequestHeader(id uint) error {
	return s.requestHeaderDAO.Delete(id)
}

// Param

func (s *RESTService) CreateRequestParam(paramRequest *OutgoingRequestParamRequest) (*OutgoingRequestParam, error) {
	param := outgoingRequestParamRequestToDBModel(paramRequest)
	err := s.requestParamDAO.Create(param)
	return param, err
}

func (s *RESTService) GetRequestParam(id uint) (*OutgoingRequestParam, error) {
	return s.requestParamDAO.Get(id)
}

func (s *RESTService) GetAllRequestParamsByConfigID(id uint) ([]*OutgoingRequestParam, error) {
	return s.requestParamDAO.GetAllByConfigID(id)
}

func (s *RESTService) UpdateRequestParam(id uint, paramRequest *OutgoingRequestParamRequest) (*OutgoingRequestParam, error) {
	param := outgoingRequestParamRequestToDBModel(paramRequest)
	param.ID = id
	err := s.requestParamDAO.Update(param)
	return param, err
}

func (s *RESTService) DeleteRequestParam(id uint) error {
	return s.requestParamDAO.Delete(id)
}
