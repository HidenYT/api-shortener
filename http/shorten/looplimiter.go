package http_shortener

import (
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

type LoopLimiterSettings struct {
	MaxRequests uint `env:"LOOP_LIMITER_MAX_REQUESTS,required,notEmpty"`
}

type ILoopLimiter interface {
	AddNewRequest(apiId uint) bool
	RemoveRequest(apiId uint)
}

type MapLoopLimiter struct {
	requests map[uint]uint
	mutex    sync.Mutex
	settings *LoopLimiterSettings
}

func (l *MapLoopLimiter) AddNewRequest(apiId uint) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.requests[apiId] == l.settings.MaxRequests {
		return false
	}
	l.requests[apiId] += 1
	return true
}

func (l *MapLoopLimiter) RemoveRequest(apiId uint) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.requests[apiId] == 0 {
		return
	}
	l.requests[apiId] -= 1
	if l.requests[apiId] == 0 {
		delete(l.requests, apiId)
	}
}

func NewLoopLimiterSettings() *LoopLimiterSettings {
	var cfg LoopLimiterSettings
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("Couldn't parse LoopLimiterSettings from env: %s", err)
	}
	return &cfg
}

func NewLoopLimiter(settings *LoopLimiterSettings) ILoopLimiter {
	return &MapLoopLimiter{
		requests: make(map[uint]uint),
		mutex:    sync.Mutex{},
		settings: settings,
	}
}
