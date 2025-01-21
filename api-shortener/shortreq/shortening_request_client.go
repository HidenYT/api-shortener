package shortreq

import (
	"io"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

type IOutgoingRequestClient interface {
	MakeRequest(request *http.Request) ([]byte, int, error)
}

type OutgoingRequestClientSettings struct {
	Timeout time.Duration `env:"OUTGOING_REQUEST_CLIENT_TIMEOUT" envDefault:"1s"`
	Retries uint8         `env:"OUTGOING_REQUEST_CLIENT_RETRIES_COUNT" envDefault:"1"`
}

type OutgoingRequestClient struct {
	client   *http.Client
	settings *OutgoingRequestClientSettings
}

func (c *OutgoingRequestClient) MakeRequest(request *http.Request) ([]byte, int, error) {
	c.client.Timeout = c.settings.Timeout
	var body []byte
	attemptsCount := 0
	statusCode := 0
	err := retry.Do(func() error {
		attemptsCount++
		response, err := c.client.Do(request)
		if err != nil {
			return err
		}
		body, err = io.ReadAll(response.Body)
		statusCode = response.StatusCode
		return err
	}, retry.Attempts(uint(c.settings.Retries)))
	logrus.Infof("Finished request to %s%s in %d attempts", request.URL.Host, request.URL.RequestURI(), attemptsCount)
	return body, statusCode, err
}

func NewOutgoingRequestClientSettings() *OutgoingRequestClientSettings {
	var cfg OutgoingRequestClientSettings
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func NewOutgoingRequestClient(settings *OutgoingRequestClientSettings) IOutgoingRequestClient {
	return &OutgoingRequestClient{
		client:   &http.Client{},
		settings: settings,
	}
}
