// http 客户端

package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"xuedengwang/core/stringx"

	"github.com/pkg/errors"
)

const (
	// ContentTypeJSON json format
	ContentTypeJSON = "application/json; charset=utf-8"
	// ContentTypeForm form format
	ContentTypeForm = "application/x-www-form-urlencoded; charset=utf-8"

	// DefaultTimeout max exec time for a request
	DefaultTimeout = 3 * time.Second
)

type Header map[string][]string

//-------------------------------------------

// Get data by get method
func Get(ctx context.Context, url string, data url.Values, options ...Option) ([]byte, error) {
	return withBody(ctx, http.MethodGet, url, []byte(data.Encode()), options...)
}

// Post get data by post method
func Post(ctx context.Context, url string, v interface{}, options ...Option) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.Wrapf(err, "[httpClient] post request struct:%v to json err", v)
	}
	return withBody(ctx, http.MethodPost, url, data, options...)
}

func withBody(ctx context.Context, method, url string, data []byte, options ...Option) (body []byte, err error) {
	opt := defaultOptions()
	for _, o := range options {
		o(opt)
	}

	retryDelay := opt.retryDelay
	if retryDelay <= 0 {
		retryDelay = DefaultRetryDelay
	}
	if opt.retryTimes == 0 {
		body, _, err = doReq(ctx, method, url, data, opt)
	} else {
		var httpCode int
		for k := 0; k < opt.retryTimes; k++ {
			body, httpCode, err = doReq(ctx, method, url, data, opt)
			if shouldRetry(ctx, httpCode) || (opt.retryVerify != nil && opt.retryVerify(body)) {
				time.Sleep(retryDelay)
				continue
			}
			return
		}
	}
	return
}

func doReq(ctx context.Context, method, url string, data []byte, opt *options) (ret []byte, code int, err error) {
	var req *http.Request
	switch method {
	case http.MethodGet:
		req, err = http.NewRequestWithContext(ctx, method, url, strings.NewReader(""))
	case http.MethodPost:
		opt.header["Content-Type"] = []string{ContentTypeJSON}
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(data))
	}

	if err != nil {
		return nil, -1, errors.Wrapf(err, "[httpClient] get req err")
	}

	client := &http.Client{
		Timeout: opt.timeout,
	}

	if opt.transport != nil {
		client.Transport = opt.transport
	}

	for key, value := range opt.header {
		req.Header.Set(key, value[0])
	}
	if method == http.MethodGet {
		req.URL.RawQuery = string(data)
	}

	resp, err := client.Do(req)
	if err != nil || resp == nil || resp.StatusCode != http.StatusOK {

		return nil, StatusDoReqErr, errors.Wrapf(err, "[httpClient] do request from [%s %s] err", method, url)
	}
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		return nil, StatusReadRespErr, errors.Wrapf(err, "[httpClient] read resp body from [%s %s] err", method, url)
	}
	return body, http.StatusOK, nil
}

// ------------------ JSON ------------------

// GetJSON get json data by get method
func GetJSON(ctx context.Context, url string, options ...Option) ([]byte, error) {
	return withJSONBody(ctx, http.MethodGet, url, nil, options...)
}

// PostJSON send json data by post method
func PostJSON(ctx context.Context, url string, data json.RawMessage, options ...Option) ([]byte, error) {
	return withJSONBody(ctx, http.MethodPost, url, data, options...)
}

func withJSONBody(ctx context.Context, method, url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	opt := defaultOptions()
	for _, o := range options {
		o(opt)
	}
	opt.header["Content-Type"] = []string{ContentTypeJSON}
	return doRequest(ctx, method, url, raw, opt)
}

// ------------------ request form ------------------

// PostForm send form data by post method
func PostForm(ctx context.Context, url string, form url.Values, options ...Option) ([]byte, error) {
	return withFormBody(ctx, http.MethodPost, url, form, options...)
}

func withFormBody(ctx context.Context, method, url string, form url.Values, options ...Option) (body []byte, err error) {
	opt := defaultOptions()
	for _, o := range options {
		o(opt)
	}
	opt.header["Content-Type"] = []string{ContentTypeForm}
	formValue := form.Encode()
	return doRequest(ctx, method, url, stringx.ToBytes(formValue), opt)
}

func doRequest(ctx context.Context, method, url string, payload []byte, opt *options) (ret []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrapf(err, "[httpClient] get req err")
	}

	client := &http.Client{
		Timeout: opt.timeout,
	}

	for key, value := range opt.header {
		req.Header.Set(key, value[0])
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "[httpClient] do request from [%s %s] err", method, url)
	}
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}

	if !isSuccess(resp.StatusCode) {
		return nil, errors.Errorf("[httpClient] status code is %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "[httpClient] read resp body from [%s %s] err", method, url)
	}
	return body, nil
}

// isSuccess check is success
func isSuccess(statusCode int) bool {
	return statusCode < http.StatusBadRequest
}
