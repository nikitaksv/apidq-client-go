package apidq

import (
	"context"
	"net/http"

	"github.com/nikitaksv/apidq-client-go/dto/address"
)

type AddressService struct {
	*service
}

// Clean Стандартизация адреса
func (s AddressService) Clean(ctx context.Context, req *address.CleanRequest) (rsp *address.CleanResponse, httpRsp *http.Response, err error) {
	rsp = &address.CleanResponse{}
	httpRsp, err = s.post(ctx, ServiceAddress, "/v1/clean/address", req, rsp)
	if err != nil {
		return
	}
	return
}

// CleanIqdq Стандартизация адреса в формате старого API
func (s AddressService) CleanIqdq(ctx context.Context, req *address.CleanRequest) (rsp *address.CleanIqdqResponse, httpRsp *http.Response, err error) {
	rsp = &address.CleanIqdqResponse{}
	httpRsp, err = s.post(ctx, ServiceAddress, "/v1/clean/address/iqdq", req, rsp)
	if err != nil {
		return
	}
	return
}

// Suggest Подсказки адреса
func (s AddressService) Suggest(ctx context.Context, req *address.SuggestRequest) (rsp *address.SuggestResponse, httpRsp *http.Response, err error) {
	rsp = &address.SuggestResponse{}
	httpRsp, err = s.post(ctx, ServiceAddress, "/v1/suggest/address", req, rsp)
	if err != nil {
		return
	}
	return
}
