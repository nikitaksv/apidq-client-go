package apidq

import (
	"context"
	"net/http"

	"github.com/nikitaksv/apidq-client-go/dto/phone"
)

type PhoneService struct {
	*service
}

// Clean Стандартизация телефонного номера
func (s PhoneService) Clean(ctx context.Context, req *phone.CleanRequest) (rsp *phone.CleanResponse, httpRsp *http.Response, err error) {
	rsp = &phone.CleanResponse{}
	httpRsp, err = s.post(ctx, ServicePhone, "/v1/clean/phone", req, rsp)
	if err != nil {
		return
	}
	return
}
