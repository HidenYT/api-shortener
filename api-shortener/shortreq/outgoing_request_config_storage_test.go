package shortreq_test

import (
	"api-shortener/shortreq"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRequestConfig(t *testing.T, apiID uint) *shortreq.OutgoingRequestConfig {
	config := &shortreq.OutgoingRequestConfig{
		Url:            "http://someurl.com/",
		Method:         "POST",
		Body:           `{"some-key":[1, 2, "3"]}`,
		ShortenedAPIID: apiID,
	}
	err := testOutgoingRequestConfigDAO.Create(config)

	require.NoError(t, err)
	require.NotEmpty(t, config)

	require.NotZero(t, config.ID)
	require.NotZero(t, config.CreatedAt)

	return config
}

func assertOutgoingRequestConfigsEqual(t *testing.T, expected, actual *shortreq.OutgoingRequestConfig) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Method, actual.Method)
	require.Equal(t, expected.Body, actual.Body)
	require.Equal(t, expected.Url, actual.Url)
	require.Equal(t, expected.ShortenedAPIID, actual.ShortenedAPIID)
}

func TestCreateRequestConfig(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	createRequestConfig(t, apiID)
}

func TestGetRequestConfig(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	config1 := createRequestConfig(t, apiID)

	config2, err := testOutgoingRequestConfigDAO.Get(config1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, config2)
	assertOutgoingRequestConfigsEqual(t, config1, config2)
}

func TestGetRequestConfigByAPIID(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	config1 := createRequestConfig(t, apiID)

	config2, err := testOutgoingRequestConfigDAO.GetByAPIID(apiID)
	require.NoError(t, err)
	require.NotEmpty(t, config2)
	assertOutgoingRequestConfigsEqual(t, config1, config2)
}

func TestUpdateRequestConfig(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	config1 := createRequestConfig(t, apiID)

	config1.Body = "aboba"
	config1.Method = "PROSTO_METHOD"
	config1.Url = "http://unknownUrl.com/"

	err := testOutgoingRequestConfigDAO.Update(config1)
	require.NoError(t, err)

	config2, err := testOutgoingRequestConfigDAO.Get(config1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, config2)
	assertOutgoingRequestConfigsEqual(t, config1, config2)
}

func TestDeleteRequestConfig(t *testing.T) {
	apiID := createShortenedAPI(t).ID
	config1 := createRequestConfig(t, apiID)

	err := testOutgoingRequestConfigDAO.Delete(config1.ID)
	require.NoError(t, err)

	config2, err := testOutgoingRequestConfigDAO.Get(config1.ID)
	require.Error(t, err)
	require.Empty(t, config2)
}
