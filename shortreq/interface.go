package shortreq

type IShortenedAPIDAO interface {
	Create() (*ShortenedAPI, error)
	Get(id uint) (*ShortenedAPI, error)
	Delete(id uint) error
}

type IOutgoingRequestConfigDAO interface {
	Create(api *OutgoingRequestConfig) error
	Get(id uint) (*OutgoingRequestConfig, error)
	GetByAPIID(apiID uint) (*OutgoingRequestConfig, error)
	Update(api *OutgoingRequestConfig) error
	Delete(id uint) error
}

type IOutgoingRequestHeaderDAO interface {
	Create(api *OutgoingRequestHeader) error
	Get(id uint) (*OutgoingRequestHeader, error)
	GetAllByConfigID(configID uint) ([]*OutgoingRequestHeader, error)
	Update(api *OutgoingRequestHeader) error
	Delete(id uint) error
}

type IOutgoingRequestParamDAO interface {
	Create(api *OutgoingRequestParam) error
	Get(id uint) (*OutgoingRequestParam, error)
	GetAllByConfigID(configID uint) ([]*OutgoingRequestParam, error)
	Update(api *OutgoingRequestParam) error
	Delete(id uint) error
}

type IShorteningRuleDAO interface {
	Create(api *ShorteningRule) error
	Get(id uint) (*ShorteningRule, error)
	GetAllByAPIID(apiID uint) ([]*ShorteningRule, error)
	Update(api *ShorteningRule) error
	Delete(id uint) error
}

type IAPIService interface {
	CreateAPI() (*ShortenedAPI, error)
	DeleteAPI(id uint) error
}
