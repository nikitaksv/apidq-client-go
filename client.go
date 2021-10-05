package apidq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BaseURL         = "https://api.apidq.io/"
	contentTypeJSON = "application/json"
	acceptJSON      = "application/json"
	authorization   = "Authorization"

	ServiceAddress = "address"
	ServicePhone   = "phone"
	ServiceName    = "name"

	ctxKeyService ctxKey = iota
)

type ctxKey int

type service struct {
	client *Client
}

func (s *service) prepareCtx(ctx context.Context, service string) context.Context {
	return context.WithValue(ctx, ctxKeyService, service)
}

func (s *service) post(ctx context.Context, service, url string, req, rsp interface{}) (*http.Response, error) {
	ctx = s.prepareCtx(ctx, service)
	r, err := s.client.newRequest(ctx, http.MethodPost, url, contentTypeJSON, req)
	if err != nil {
		return nil, err
	}

	httpRsp, err := s.client.do(ctx, r, rsp)
	if err != nil {
		return httpRsp, err
	}
	return httpRsp, nil
}

type RequestOptionFunc func(r *http.Request) error

type Client struct {
	baseURL *url.URL
	client  *http.Client

	common  service
	Address *AddressService
	Phone   *PhoneService
	Name    *NameService

	requestOptions []RequestOptionFunc
}

func NewClient(httpClient *http.Client, baseURL string, reqOpts ...RequestOptionFunc) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   15 * time.Second,
		}
	}
	if baseURL == "" {
		baseURL = BaseURL
	}

	pBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, baseURL: pBaseURL, requestOptions: reqOpts}

	c.common.client = c
	c.Address = &AddressService{&c.common}
	c.Phone = &PhoneService{&c.common}
	c.Name = &NameService{&c.common}

	return c, nil
}

// WithAuthService ApiDQ дает возможность генерировать отдельные токены для всех сервисов.
// Допустимые значения service = ["address","phone","name"]
func (c *Client) WithAuthService(apiKey, service string) *Client {
	c.requestOptions = append(c.requestOptions, func(r *http.Request) error {
		if strings.Contains(fmt.Sprintf("%v", r.Context().Value(ctxKeyService)), service) {
			r.Header.Set(authorization, apiKey)
		}
		return nil
	})
	return c
}

// WithAuth Auth for all services
func (c *Client) WithAuth(apiKey string) *Client {
	c.requestOptions = append(c.requestOptions, func(r *http.Request) error {
		r.Header.Set(authorization, apiKey)
		return nil
	})
	return c
}

// WithReqOptions Add request options
func (c *Client) WithReqOptions(reqOpts ...RequestOptionFunc) *Client {
	c.requestOptions = append(c.requestOptions, reqOpts...)
	return c
}

func (c *Client) newRequest(ctx context.Context, method, url, contentType string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		if method == http.MethodPost {
			buf = &bytes.Buffer{}
			enc := json.NewEncoder(buf)
			errEnc := enc.Encode(body)
			if errEnc != nil {
				return nil, errEnc
			}
		} else {
			return nil, fmt.Errorf("request Method \"%q\" is unknown", method)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Accept", acceptJSON)
	}

	for _, opt := range c.requestOptions {
		if errOpt := opt(req); errOpt != nil {
			return nil, errOpt
		}
	}

	return req, nil
}

// Отправка запроса
func (c *Client) do(_ context.Context, req *http.Request, v interface{}) (rsp *http.Response, err error) {
	rsp, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rsp != nil && rsp.Body != nil {
			if e := rsp.Body.Close(); e != nil && err == nil {
				err = e // if body not close, return err
			}
		}
	}()

	if v != nil && rsp.ContentLength != 0 {
		body, e := ioutil.ReadAll(rsp.Body)
		if e != nil {
			return nil, e
		}

		if rsp.StatusCode != 200 {
			errRsp := &ErrorResponse{}
			decErr := json.Unmarshal(body, errRsp)
			if decErr == nil {
				return nil, errRsp
			}
			return nil, decErr
		}

		decErr := json.Unmarshal(body, v)
		if decErr == nil {
			return rsp, nil
		}

		return nil, decErr
	}

	return rsp, nil
}

type ErrorResponse struct {
	Message string `json:"message"`
	// https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	// https://developers.google.com/maps-booking/reference/grpc-api/status_codes
	Code int `json:"code"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}
