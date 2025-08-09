package shortreq_test

import (
	"api-shortener/shortreq"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func createShortenedAPI(t *testing.T) *shortreq.ShortenedAPI {
	shortenedAPI, err := testShortenedAPIDAO.Create()

	require.NoError(t, err)
	require.NotEmpty(t, shortenedAPI)

	require.NotZero(t, shortenedAPI.ID)
	require.NotZero(t, shortenedAPI.CreatedAt)

	return shortenedAPI
}

func TestCreateShortenedAPI(t *testing.T) {
	createShortenedAPI(t)
}

func TestGetShortenedAPI(t *testing.T) {
	shortenedAPI1 := createShortenedAPI(t)

	shortenedAPI2, err := testShortenedAPIDAO.Get(shortenedAPI1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, shortenedAPI2)

	require.NotZero(t, shortenedAPI2.ID)
	require.NotZero(t, shortenedAPI2.CreatedAt)

	require.Equal(t, shortenedAPI1.ID, shortenedAPI2.ID)
}

func TestDeleteShortenedAPI(t *testing.T) {
	shortenedAPI1 := createShortenedAPI(t)

	err := testShortenedAPIDAO.Delete(shortenedAPI1.ID)
	require.NoError(t, err)

	shortenedAPI2, err := testShortenedAPIDAO.Get(shortenedAPI1.ID)
	require.Error(t, err)
	require.Empty(t, shortenedAPI2)
}

func TestGetShortenedAPIWithRules(t *testing.T) {
	shortenedAPI1 := createShortenedAPI(t)
	rules1 := []*shortreq.ShorteningRule{createShorteningRule(t, shortenedAPI1.ID), createShorteningRule(t, shortenedAPI1.ID)}

	shortenedAPI2, err := testShortenedAPIDAO.Get(shortenedAPI1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, shortenedAPI2)

	require.NotZero(t, shortenedAPI2.ID)
	require.NotZero(t, shortenedAPI2.CreatedAt)

	rules2 := shortenedAPI2.ShorteningRules

	require.Equal(t, shortenedAPI1.ID, shortenedAPI2.ID)
	require.Equal(t, len(rules1), len(rules2))
	sort.Slice(rules1, func(i, j int) bool { return rules1[i].ID < rules1[j].ID })
	sort.Slice(rules2, func(i, j int) bool { return rules2[i].ID < rules2[j].ID })
	for i := range rules1 {
		assertShorteningRulesEqual(t, rules1[i], &rules2[i])
	}
}
