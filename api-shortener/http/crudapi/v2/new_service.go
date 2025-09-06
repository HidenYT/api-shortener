package http

import (
	"api-shortener/shortreq"
)

type IAPIDTOService interface {
	Create(*ShortenedAPIDTO) (*ShortenedAPIDTO, error)
	GetByID(id uint) (*ShortenedAPIDTO, error)
	Update(uint, *ShortenedAPIDTO) (*ShortenedAPIDTO, error)
	DeleteByID(id uint) error
}

type APIDTOService struct {
	apiDAO           shortreq.IShortenedAPIDAO
	requestConfigDAO shortreq.IOutgoingRequestConfigDAO
	ruleDAO          shortreq.IShorteningRuleDAO
	requestHeaderDAO shortreq.IOutgoingRequestHeaderDAO
	requestParamDAO  shortreq.IOutgoingRequestParamDAO
}

func (s APIDTOService) createOutgoingRequestConfig(
	apiID uint, dto *OutgoingRequestConfigDTO,
) (*shortreq.OutgoingRequestConfig, error) {
	requestConfigDBModel := &shortreq.OutgoingRequestConfig{
		Url:            dto.Url,
		Method:         dto.Method,
		Body:           dto.Body,
		ShortenedAPIID: apiID,
	}
	err := s.requestConfigDAO.Create(requestConfigDBModel)
	return requestConfigDBModel, err
}

func (s APIDTOService) createRules(dto *ShortenedAPIDTO) error {
	for _, rule := range dto.ShorteningRules {
		ruleDBModel := &shortreq.ShorteningRule{
			FieldName:       rule.FieldName,
			FieldValueQuery: rule.FieldValueQuery,
			ShortenedAPIID:  dto.ID,
		}
		if err := s.ruleDAO.Create(ruleDBModel); err != nil {
			return err
		}
		rule.ID = ruleDBModel.ID
	}
	return nil
}

func (s APIDTOService) createHeaders(dto *OutgoingRequestConfigDTO) error {
	for _, header := range dto.Headers {
		headerDBModel := &shortreq.OutgoingRequestHeader{
			Name:                    header.Name,
			Value:                   header.Value,
			OutgoingRequestConfigID: dto.ID,
		}
		if err := s.requestHeaderDAO.Create(headerDBModel); err != nil {
			return err
		}
		header.ID = headerDBModel.ID
	}
	return nil
}

func (s APIDTOService) createParams(dto *OutgoingRequestConfigDTO) error {
	for _, param := range dto.Params {
		paramDBModel := &shortreq.OutgoingRequestParam{
			Name:                    param.Name,
			Value:                   param.Value,
			OutgoingRequestConfigID: dto.ID,
		}
		if err := s.requestParamDAO.Create(paramDBModel); err != nil {
			return err
		}
		param.ID = paramDBModel.ID
	}
	return nil
}

func (s APIDTOService) Create(dto *ShortenedAPIDTO) (*ShortenedAPIDTO, error) {
	apiDBModel, err := s.apiDAO.Create()
	if err != nil {
		return nil, err
	}
	dto.ID = apiDBModel.ID

	requestConfigDBModel, err := s.createOutgoingRequestConfig(apiDBModel.ID, &dto.OutgoingRequestConfig)
	if err != nil {
		return nil, err
	}
	dto.OutgoingRequestConfig.ID = requestConfigDBModel.ID

	if err := s.createRules(dto); err != nil {
		return nil, err
	}
	if err := s.createHeaders(&dto.OutgoingRequestConfig); err != nil {
		return nil, err
	}
	if err := s.createParams(&dto.OutgoingRequestConfig); err != nil {
		return nil, err
	}

	return dto, nil
}

func (s APIDTOService) GetByID(id uint) (*ShortenedAPIDTO, error) {
	apiDBModel, err := s.apiDAO.Get(id)
	if err != nil {
		return nil, err
	}
	outgoingRequestConfigDBModel, err := s.requestConfigDAO.GetByAPIID(id)
	if err != nil {
		return nil, err
	}
	rules, err := s.ruleDAO.GetAllByAPIID(id)
	if err != nil {
		return nil, err
	}
	headers, err := s.requestHeaderDAO.GetAllByConfigID(outgoingRequestConfigDBModel.ID)
	if err != nil {
		return nil, err
	}
	params, err := s.requestParamDAO.GetAllByConfigID(outgoingRequestConfigDBModel.ID)
	if err != nil {
		return nil, err
	}
	headersDTO := make([]*OutgoingRequestHeaderDTO, len(headers))
	for i, h := range headers {
		headersDTO[i] = &OutgoingRequestHeaderDTO{
			ID:    h.ID,
			Name:  h.Name,
			Value: h.Value,
		}
	}
	paramsDTO := make([]*OutgoingRequestParamDTO, len(params))
	for i, p := range params {
		paramsDTO[i] = &OutgoingRequestParamDTO{
			ID:    p.ID,
			Name:  p.Name,
			Value: p.Value,
		}
	}
	rulesDTO := make([]*ShorteningRuleDTO, len(rules))
	for i, r := range rules {
		rulesDTO[i] = &ShorteningRuleDTO{
			ID:              r.ID,
			FieldName:       r.FieldName,
			FieldValueQuery: r.FieldValueQuery,
		}
	}
	return &ShortenedAPIDTO{
		ID: apiDBModel.ID,
		OutgoingRequestConfig: OutgoingRequestConfigDTO{
			ID:      outgoingRequestConfigDBModel.ID,
			Url:     outgoingRequestConfigDBModel.Url,
			Method:  outgoingRequestConfigDBModel.Method,
			Body:    outgoingRequestConfigDBModel.Body,
			Headers: headersDTO,
			Params:  paramsDTO,
		},
		ShorteningRules: rulesDTO,
	}, nil
}

func (s APIDTOService) Delete(id uint) error {
	return s.apiDAO.Delete(id)
}

func splitToCreateUpdateDeleteNamedEntities[T interface {
	GetName() string
	GetID() uint
	SetID(uint)
}](existing, got []T) ([]T, []T, []T) {
	create := make([]T, 0)
	update := make([]T, 0)
	delete := make([]T, 0)
	for _, gt := range got {
		foundExisting := false
		for _, ex := range existing {
			if gt.GetName() == ex.GetName() {
				gt.SetID(ex.GetID())
				update = append(update, gt)
				foundExisting = true
				break
			}
		}
		if !foundExisting {
			create = append(create, gt)
		}
	}
	for _, ex := range existing {
		found := false
		for _, gt := range got {
			if gt.GetName() == ex.GetName() {
				found = true
				break
			}
		}
		if !found {
			delete = append(delete, ex)
		}
	}
	return create, update, delete
}

func (s APIDTOService) updateRequestConfig(dto *OutgoingRequestConfigDTO, dbModel *shortreq.OutgoingRequestConfig) error {
	dto.ID = dbModel.ID
	dbModel.Body = dto.Body
	dbModel.Url = dto.Url
	dbModel.Method = dto.Method
	return s.requestConfigDAO.Update(dbModel)
}

func (s APIDTOService) makeUpdatedRules(apiID uint, gotRules []*ShorteningRuleDTO) ([]*ShorteningRuleDTO, error) {
	existingRules, err := s.ruleDAO.GetAllByAPIID(apiID)
	if err != nil {
		return []*ShorteningRuleDTO{}, err
	}
	gotDBRules := make([]*shortreq.ShorteningRule, len(gotRules))
	for i, r := range gotRules {
		dbRule := shortreq.ShorteningRule{
			FieldName:       r.FieldName,
			FieldValueQuery: r.FieldValueQuery,
			ShortenedAPIID:  apiID,
		}
		gotDBRules[i] = &dbRule
	}
	create, update, delete := splitToCreateUpdateDeleteNamedEntities(existingRules, gotDBRules)
	result := make([]*ShorteningRuleDTO, 0)
	for _, r := range create {
		if err := s.ruleDAO.Create(r); err != nil {
			return []*ShorteningRuleDTO{}, err
		}
		result = append(result, &ShorteningRuleDTO{ID: r.ID, FieldName: r.FieldName, FieldValueQuery: r.FieldValueQuery})
	}
	for _, r := range update {
		if err := s.ruleDAO.Update(r); err != nil {
			return []*ShorteningRuleDTO{}, err
		}
		result = append(result, &ShorteningRuleDTO{ID: r.ID, FieldName: r.FieldName, FieldValueQuery: r.FieldValueQuery})
	}
	for _, r := range delete {
		if err := s.ruleDAO.Delete(r.ID); err != nil {
			return []*ShorteningRuleDTO{}, err
		}
	}
	return result, nil
}

func (s APIDTOService) makeUpdatedHeaders(configID uint, gotHeaders []*OutgoingRequestHeaderDTO) ([]*OutgoingRequestHeaderDTO, error) {
	existingHeaders, err := s.requestHeaderDAO.GetAllByConfigID(configID)
	if err != nil {
		return []*OutgoingRequestHeaderDTO{}, err
	}
	gotDBHeaders := make([]*shortreq.OutgoingRequestHeader, len(gotHeaders))
	for i, r := range gotHeaders {
		dbHeader := shortreq.OutgoingRequestHeader{
			Name:                    r.Name,
			Value:                   r.Value,
			OutgoingRequestConfigID: configID,
		}
		gotDBHeaders[i] = &dbHeader
	}
	create, update, delete := splitToCreateUpdateDeleteNamedEntities(existingHeaders, gotDBHeaders)
	result := make([]*OutgoingRequestHeaderDTO, 0)
	for _, r := range create {
		if err := s.requestHeaderDAO.Create(r); err != nil {
			return []*OutgoingRequestHeaderDTO{}, err
		}
		result = append(result, &OutgoingRequestHeaderDTO{ID: r.ID, Name: r.Name, Value: r.Value})
	}
	for _, r := range update {
		if err := s.requestHeaderDAO.Update(r); err != nil {
			return []*OutgoingRequestHeaderDTO{}, err
		}
		result = append(result, &OutgoingRequestHeaderDTO{ID: r.ID, Name: r.Name, Value: r.Value})
	}
	for _, r := range delete {
		if err := s.requestHeaderDAO.Delete(r.ID); err != nil {
			return []*OutgoingRequestHeaderDTO{}, err
		}
	}
	return result, nil
}

func (s APIDTOService) makeUpdatedParams(configID uint, gotParams []*OutgoingRequestParamDTO) ([]*OutgoingRequestParamDTO, error) {
	existingParams, err := s.requestParamDAO.GetAllByConfigID(configID)
	if err != nil {
		return []*OutgoingRequestParamDTO{}, err
	}
	gotDBParams := make([]*shortreq.OutgoingRequestParam, len(gotParams))
	for i, r := range gotParams {
		dbParam := shortreq.OutgoingRequestParam{
			Name:                    r.Name,
			Value:                   r.Value,
			OutgoingRequestConfigID: configID,
		}
		gotDBParams[i] = &dbParam
	}
	create, update, delete := splitToCreateUpdateDeleteNamedEntities(existingParams, gotDBParams)
	result := make([]*OutgoingRequestParamDTO, 0)
	for _, r := range create {
		if err := s.requestParamDAO.Create(r); err != nil {
			return []*OutgoingRequestParamDTO{}, err
		}
		result = append(result, &OutgoingRequestParamDTO{ID: r.ID, Name: r.Name, Value: r.Value})
	}
	for _, r := range update {
		if err := s.requestParamDAO.Update(r); err != nil {
			return []*OutgoingRequestParamDTO{}, err
		}
		result = append(result, &OutgoingRequestParamDTO{ID: r.ID, Name: r.Name, Value: r.Value})
	}
	for _, r := range delete {
		if err := s.requestParamDAO.Delete(r.ID); err != nil {
			return []*OutgoingRequestParamDTO{}, err
		}
	}
	return result, nil
}

func (s APIDTOService) Update(apiID uint, dto *ShortenedAPIDTO) (*ShortenedAPIDTO, error) {
	requestConfigDBModel, err := s.requestConfigDAO.GetByAPIID(apiID)
	if err != nil {
		return nil, err
	}

	dto.ID = apiID
	if err = s.updateRequestConfig(&dto.OutgoingRequestConfig, requestConfigDBModel); err != nil {
		return nil, err
	}

	updatedRules, err := s.makeUpdatedRules(dto.ID, dto.ShorteningRules)
	if err != nil {
		return nil, err
	}
	dto.ShorteningRules = updatedRules

	updatedHeaders, err := s.makeUpdatedHeaders(requestConfigDBModel.ID, dto.OutgoingRequestConfig.Headers)
	if err != nil {
		return nil, err
	}
	dto.OutgoingRequestConfig.Headers = updatedHeaders

	updatedParams, err := s.makeUpdatedParams(requestConfigDBModel.ID, dto.OutgoingRequestConfig.Params)
	if err != nil {
		return nil, err
	}
	dto.OutgoingRequestConfig.Params = updatedParams
	return dto, nil
}

func (s APIDTOService) DeleteByID(apiID uint) error {
	return s.apiDAO.Delete(apiID)
}

func NewAPIDTOService(
	apiDAO shortreq.IShortenedAPIDAO,
	requestConfigDAO shortreq.IOutgoingRequestConfigDAO,
	ruleDAO shortreq.IShorteningRuleDAO,
	requestHeaderDAO shortreq.IOutgoingRequestHeaderDAO,
	requestParamDAO shortreq.IOutgoingRequestParamDAO,
) APIDTOService {
	return APIDTOService{
		apiDAO:           apiDAO,
		requestConfigDAO: requestConfigDAO,
		ruleDAO:          ruleDAO,
		requestHeaderDAO: requestHeaderDAO,
		requestParamDAO:  requestParamDAO,
	}
}
