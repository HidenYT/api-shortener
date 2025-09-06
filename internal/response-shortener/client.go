package shortener

import (
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

type OutgoingRequestClientSettings struct {
	Timeout time.Duration `env:"OUTGOING_REQUEST_CLIENT_TIMEOUT" envDefault:"1s"`
	Retries uint8         `env:"OUTGOING_REQUEST_CLIENT_RETRIES_COUNT" envDefault:"1"`
}

type OutgoingRequestClient struct {
	client   *http.Client
	settings *OutgoingRequestClientSettings
}

func (c *OutgoingRequestClient) MakeRequest(request *http.Request) (*http.Response, error) {
	c.client.Timeout = c.settings.Timeout
	var response *http.Response
	attemptsCount := 0
	err := retry.Do(func() error {
		attemptsCount++
		var err error
		response, err = c.client.Do(request)
		return err
	}, retry.Attempts(uint(c.settings.Retries)))
	logrus.Infof("Finished request to %s://%s%s in %d attempts", request.URL.Scheme, request.URL.Host, request.URL.RequestURI(), attemptsCount)
	return response, err
}

func NewOutgoingRequestClientSettings() *OutgoingRequestClientSettings {
	var cfg OutgoingRequestClientSettings
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("Couldn't parse OutgoingRequestClientSettings from env: %s", err)
	}
	return &cfg
}

func NewOutgoingRequestClient(settings *OutgoingRequestClientSettings) IOutgoingRequestClient {
	return &OutgoingRequestClient{
		client:   &http.Client{},
		settings: settings,
	}
}
