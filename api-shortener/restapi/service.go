package restapi

import (
	"api-shortener/shortreq"

	"gorm.io/gorm"
)

type IRESTService interface {
	Create(api *shortreq.ShortenedAPI) (*shortreq.ShortenedAPI, error)
	Get(id uint) (*shortreq.ShortenedAPI, error)
	Update(api *shortreq.ShortenedAPI) (*shortreq.ShortenedAPI, error)
	Delete(id uint) error
}

type RESTService struct {
	db  *gorm.DB
	dao shortreq.IShortenedAPIDAO
}

func NewRESTService(db *gorm.DB, dao shortreq.IShortenedAPIDAO) IRESTService {
	return &RESTService{db: db, dao: dao}
}

func (s *RESTService) Create(api *shortreq.ShortenedAPI) (*shortreq.ShortenedAPI, error) {
	if api.Config.Headers == nil {
		api.Config.Headers = []*shortreq.OutgoingRequestHeader{}
	}
	if api.Config.Params == nil {
		api.Config.Params = []*shortreq.OutgoingRequestParam{}
	}
	err := s.dao.Create(api)
	return api, err
}

func (s *RESTService) Get(id uint) (*shortreq.ShortenedAPI, error) {
	return s.dao.Get(id)
}

func (s *RESTService) Update(api *shortreq.ShortenedAPI) (*shortreq.ShortenedAPI, error) {
	if api.Config.Headers == nil {
		api.Config.Headers = []*shortreq.OutgoingRequestHeader{}
	}
	if api.Config.Params == nil {
		api.Config.Params = []*shortreq.OutgoingRequestParam{}
	}
	err := s.dao.Update(api)
	return api, err
}

func (s *RESTService) Delete(id uint) error {
	return s.dao.Delete(id)
}
