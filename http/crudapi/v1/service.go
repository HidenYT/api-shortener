package http

import (
	"github.com/HidenYT/api-shortener/shortreq"
)

// API

type APIService struct {
	apiDAO shortreq.IShortenedAPIDAO
}

func NewAPIService(apiDAO shortreq.IShortenedAPIDAO) *APIService {
	return &APIService{apiDAO: apiDAO}
}

func (s *APIService) Create() (*shortreq.ShortenedAPI, error) {
	return s.apiDAO.Create()
}

func (s *APIService) GetByID(id uint) (*shortreq.ShortenedAPI, error) {
	return s.apiDAO.Get(id)
}

func (s *APIService) Delete(id uint) error {
	return s.apiDAO.Delete(id)
}

// RequestConfig

type RequestConfigService struct {
	requestConfigDAO shortreq.IOutgoingRequestConfigDAO
}

func NewRequestConfigService(configDAO shortreq.IOutgoingRequestConfigDAO) *RequestConfigService {
	return &RequestConfigService{requestConfigDAO: configDAO}
}

func (s *RequestConfigService) Create(configRequest *OutgoingRequestConfigRequest) (*shortreq.OutgoingRequestConfig, error) {
	config := outgoingRequestConfigRequestToDBModel(configRequest)
	err := s.requestConfigDAO.Create(config)
	return config, err
}

func (s *RequestConfigService) GetByID(id uint) (*shortreq.OutgoingRequestConfig, error) {
	return s.requestConfigDAO.Get(id)
}

func (s *RequestConfigService) GetByAPIID(id uint) (*shortreq.OutgoingRequestConfig, error) {
	return s.requestConfigDAO.GetByAPIID(id)
}

func (s *RequestConfigService) Update(id uint, configRequest *OutgoingRequestConfigRequest) (*shortreq.OutgoingRequestConfig, error) {
	config := outgoingRequestConfigRequestToDBModel(configRequest)
	config.ID = id
	err := s.requestConfigDAO.Update(config)
	return config, err
}

func (s *RequestConfigService) Delete(id uint) error {
	return s.requestConfigDAO.Delete(id)
}

// ShorteningRule

type ShorteningRuleService struct {
	ruleDAO shortreq.IShorteningRuleDAO
}

func NewShorteningRuleService(ruleDAO shortreq.IShorteningRuleDAO) *ShorteningRuleService {
	return &ShorteningRuleService{ruleDAO: ruleDAO}
}

func (s *ShorteningRuleService) Create(ruleRequest *ShorteningRuleRequest) (*shortreq.ShorteningRule, error) {
	rule := shorteningRuleRequestToDBModel(ruleRequest)
	err := s.ruleDAO.Create(rule)
	return rule, err
}

func (s *ShorteningRuleService) GetByID(id uint) (*shortreq.ShorteningRule, error) {
	return s.ruleDAO.Get(id)
}

func (s *ShorteningRuleService) GetAllByAPIID(id uint) ([]*shortreq.ShorteningRule, error) {
	return s.ruleDAO.GetAllByAPIID(id)
}

func (s *ShorteningRuleService) Update(id uint, ruleRequest *ShorteningRuleRequest) (*shortreq.ShorteningRule, error) {
	rule := shorteningRuleRequestToDBModel(ruleRequest)
	rule.ID = id
	err := s.ruleDAO.Update(rule)
	return rule, err
}

func (s *ShorteningRuleService) Delete(id uint) error {
	return s.ruleDAO.Delete(id)
}

// Header

type RequestHeaderService struct {
	requestHeaderDAO shortreq.IOutgoingRequestHeaderDAO
}

func NewRequestHeaderService(headerDAO shortreq.IOutgoingRequestHeaderDAO) *RequestHeaderService {
	return &RequestHeaderService{requestHeaderDAO: headerDAO}
}

func (s *RequestHeaderService) Create(headerRequest *OutgoingRequestHeaderRequest) (*shortreq.OutgoingRequestHeader, error) {
	header := outgoingRequestHeaderRequestToDBModel(headerRequest)
	err := s.requestHeaderDAO.Create(header)
	return header, err
}

func (s *RequestHeaderService) GetByID(id uint) (*shortreq.OutgoingRequestHeader, error) {
	return s.requestHeaderDAO.Get(id)
}

func (s *RequestHeaderService) GetAllByConfigID(id uint) ([]*shortreq.OutgoingRequestHeader, error) {
	return s.requestHeaderDAO.GetAllByConfigID(id)
}

func (s *RequestHeaderService) Update(id uint, headerRequest *OutgoingRequestHeaderRequest) (*shortreq.OutgoingRequestHeader, error) {
	header := outgoingRequestHeaderRequestToDBModel(headerRequest)
	header.ID = id
	err := s.requestHeaderDAO.Update(header)
	return header, err
}

func (s *RequestHeaderService) Delete(id uint) error {
	return s.requestHeaderDAO.Delete(id)
}

// Param

type RequestParamService struct {
	requestParamDAO shortreq.IOutgoingRequestParamDAO
}

func NewRequestParamService(paramDAO shortreq.IOutgoingRequestParamDAO) *RequestParamService {
	return &RequestParamService{requestParamDAO: paramDAO}
}

func (s *RequestParamService) Create(paramRequest *OutgoingRequestParamRequest) (*shortreq.OutgoingRequestParam, error) {
	param := outgoingRequestParamRequestToDBModel(paramRequest)
	err := s.requestParamDAO.Create(param)
	return param, err
}

func (s *RequestParamService) GetByID(id uint) (*shortreq.OutgoingRequestParam, error) {
	return s.requestParamDAO.Get(id)
}

func (s *RequestParamService) GetAllByConfigID(id uint) ([]*shortreq.OutgoingRequestParam, error) {
	return s.requestParamDAO.GetAllByConfigID(id)
}

func (s *RequestParamService) Update(id uint, paramRequest *OutgoingRequestParamRequest) (*shortreq.OutgoingRequestParam, error) {
	param := outgoingRequestParamRequestToDBModel(paramRequest)
	param.ID = id
	err := s.requestParamDAO.Update(param)
	return param, err
}

func (s *RequestParamService) Delete(id uint) error {
	return s.requestParamDAO.Delete(id)
}
