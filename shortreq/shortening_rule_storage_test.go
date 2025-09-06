package shortreq_test

import (
	"sort"
	"testing"

	"github.com/HidenYT/api-shortener/shortreq"

	"github.com/stretchr/testify/require"
)

func createShorteningRule(t *testing.T, ShortenedAPIID uint) *shortreq.ShorteningRule {
	rule := &shortreq.ShorteningRule{
		FieldName:       "somefield",
		FieldValueQuery: "query",
		ShortenedAPIID:  ShortenedAPIID,
	}
	err := testShorteningRuleDAO.Create(rule)

	require.NoError(t, err)
	require.NotEmpty(t, rule)

	require.NotZero(t, rule.ID)
	require.NotZero(t, rule.CreatedAt)

	return rule
}

func assertShorteningRulesEqual(t *testing.T, rule1, rule2 *shortreq.ShorteningRule) {
	require.Equal(t, rule1.ID, rule2.ID)
	require.Equal(t, rule1.FieldName, rule2.FieldName)
	require.Equal(t, rule1.FieldValueQuery, rule2.FieldValueQuery)
	require.Equal(t, rule1.ShortenedAPIID, rule2.ShortenedAPIID)
}

func TestCreateShorteningRule(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	createShorteningRule(t, apiID)
}

func TestCantCreateShorteningRuleWithBadQuery(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	rule := &shortreq.ShorteningRule{
		FieldName:       "somefield",
		FieldValueQuery: "..----....",
		ShortenedAPIID:  apiID,
	}
	err := testShorteningRuleDAO.Create(rule)
	require.Error(t, err)
	require.Zero(t, rule.ID)
}

func TestGetShorteningRule(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	rule1 := createShorteningRule(t, apiID)

	rule2, err := testShorteningRuleDAO.Get(rule1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, rule2)
	assertShorteningRulesEqual(t, rule1, rule2)
}

func TestGetAllShorteningRulesByAPIID(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	rules1 := []*shortreq.ShorteningRule{createShorteningRule(t, apiID), createShorteningRule(t, apiID)}

	rules2, err := testShorteningRuleDAO.GetAllByAPIID(apiID)

	require.NoError(t, err)

	require.Equal(t, len(rules1), len(rules2))
	sort.Slice(rules1, func(i, j int) bool { return rules1[i].ID < rules1[j].ID })
	sort.Slice(rules2, func(i, j int) bool { return rules2[i].ID < rules2[j].ID })

	for i := range rules1 {
		rule1 := rules1[i]
		rule2 := rules2[i]
		require.Equal(t, rule1.ID, rule2.ID)
		require.Equal(t, rule1.FieldName, rule2.FieldName)
		require.Equal(t, rule1.FieldValueQuery, rule2.FieldValueQuery)
	}
}

func TestUpdateShorteningRule(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	rule1 := createShorteningRule(t, apiID)
	rule1.FieldName = "some-random-string"
	rule1.FieldValueQuery = "some_random_query.0.134"
	err := testShorteningRuleDAO.Update(rule1)

	require.NoError(t, err)

	rule2, err := testShorteningRuleDAO.Get(rule1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, rule2)
	assertShorteningRulesEqual(t, rule1, rule2)
}

func TestDeleteShorteningRule(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	rule1 := createShorteningRule(t, apiID)

	err := testShorteningRuleDAO.Delete(rule1.ID)
	require.NoError(t, err)

	rule2, err := testShorteningRuleDAO.Get(rule1.ID)
	require.Error(t, err)
	require.Empty(t, rule2)
}
