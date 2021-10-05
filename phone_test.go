package apidq

import (
	"context"
	"net/http"
	"testing"

	"github.com/nikitaksv/apidq-client-go/dto/phone"
)

func TestPhoneClean(t *testing.T) {
	reqBs := []byte(`{"query":"89611122333","countryCode":"RU"}`)
	rspBs := []byte(`{"original":"89611122333","international":"+7 961 112-23-33","national":"8 (961) 112-23-33","E164":"+79611122333","RFC3966":"tel:+7-961-112-23-33","carrier":"Beeline","countryCode":7,"country":"RU","areaCode":"","timezones":["Europe/Moscow"],"geocoding":"","subscriberNumber":"9611122333","type":"MOBILE","possible":true,"valid":true}`)
	testEndpointCall(t, reqBs, rspBs, func(client *Client) (interface{}, *http.Response, error) {
		return client.Phone.Clean(context.Background(), &phone.CleanRequest{
			Query:       "89611122333",
			CountryCode: "RU",
		})
	})
}
