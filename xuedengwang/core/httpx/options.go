package httpx

import (
	"net/http"
	"net/url"
	"time"
)

// Option is a function that sets some option on the client.
type Option func(c *options)

// Options control behavior of the client.
type options struct {
	header map[string][]string
	// timeout of per request
	timeout time.Duration

	retryTimes  int
	retryDelay  time.Duration
	retryVerify RetryVerify
	transport   *http.Transport
}

func defaultOptions() *options {
	return &options{
		header:     make(map[string][]string),
		timeout:    DefaultTimeout,
		retryTimes: DefaultRetryTimes,
	}
}

// WithTimeout with a timeout for per request
func WithTimeout(duration time.Duration) Option {
	return func(cfg *options) {
		cfg.timeout = duration
	}
}

func WithHeader(header map[string][]string) Option {
	return func(cfg *options) {
		cfg.header = header
	}
}

func WithTransport(proxyUrl string) Option {
	return func(cfg *options) {
		proxy := func(*http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		cfg.transport = &http.Transport{Proxy: proxy}
	}
}

func WithRetryTime(retryTime int) Option {
	return func(cfg *options) {
		cfg.retryTimes = retryTime
	}
}
