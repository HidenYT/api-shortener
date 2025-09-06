package shortener_test

import (
	"net/http"
	"net/url"
	"testing"

	shortener "github.com/HidenYT/api-shortener/internal/response-shortener"

	"github.com/stretchr/testify/require"
)

func TestShortenJSON(t *testing.T) {
	result, err := shortener.ShortenJSONFunc(testData, testRules)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, expectedResponse, result)
}

func TestShortenRawBody(t *testing.T) {
	result, err := shortener.ShortenRawBodyFunc(testDataRaw, testRules)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, expectedResponse, result)
}

func TestProcessRequest(t *testing.T) {
	settings := shortener.NewOutgoingRequestClientSettings()
	client := shortener.NewOutgoingRequestClient(settings)
	shortenerSvc := shortener.NewResponseShortener(client)

	expected := shortener.ShortenedResponse{
		JSON:       &expectedResponse,
		StatusCode: 200,
	}

	serverURL, _ := url.Parse(testHttpServer.URL)
	response, err := shortenerSvc.ProcessRequest(&http.Request{URL: serverURL}, testRules)

	require.NoError(t, err)
	require.NotEmpty(t, response)

	require.Equal(t, expected.JSON, response.JSON)
	require.Equal(t, expected.StatusCode, response.StatusCode)
	require.Equal(t, "application/json", response.Headers.Get("Content-Type"))
	require.Equal(t, "abcd123", response.Headers.Get("X-Some-Random-Header"))
}
