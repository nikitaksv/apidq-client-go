package apidq

import (
	"context"
	"net/http"

	"github.com/nikitaksv/apidq-client-go/dto/name"
)

type NameService struct {
	*service
}

// Clean Стандартизация ФИО
func (s NameService) Clean(ctx context.Context, req *name.CleanRequest) (rsp *name.CleanResponse, httpRsp *http.Response, err error) {
	rsp = &name.CleanResponse{}
	httpRsp, err = s.post(ctx, ServiceName, "/v1/clean/name", req, rsp)
	if err != nil {
		return
	}
	return
}

// Suggest Подсказки ФИО
func (s NameService) Suggest(ctx context.Context, req *name.SuggestRequest) (rsp *name.SuggestResponse, httpRsp *http.Response, err error) {
	rsp = &name.SuggestResponse{}
	httpRsp, err = s.post(ctx, ServiceName, "/v1/suggest/name", req, rsp)
	if err != nil {
		return
	}
	return
}
