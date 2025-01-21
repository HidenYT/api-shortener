package shortreq

import (
	"fmt"
)

type ShorteningRules map[string]string

type IRulesResolver interface {
	GetRules(api *ShortenedAPI) (ShorteningRules, error)
}

type ShorteningRulesResolver struct {
	repo IShortenedAPIDAO
}

type ShorteningRulesNotFound struct {
	Id uint
}

func (e *ShorteningRulesNotFound) Error() string {
	return fmt.Sprintf("Rules with ID %d not found", e.Id)
}

func (r *ShorteningRulesResolver) GetRules(api *ShortenedAPI) (ShorteningRules, error) {
	resultRules := make(map[string]string)
	for _, rule := range api.ShorteningRules {
		resultRules[rule.FieldName] = rule.FieldValueQuery
	}
	return resultRules, nil
}

func NewShorteningRulesResolver(repo IShortenedAPIDAO) IRulesResolver {
	return &ShorteningRulesResolver{repo: repo}
}
