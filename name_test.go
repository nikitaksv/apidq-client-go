package apidq

import (
	"context"
	"net/http"
	"testing"

	"github.com/nikitaksv/apidq-client-go/dto/name"
)

func TestNameClean(t *testing.T) {
	reqBs := []byte(`{"query":"Андрей Ильич Петров"}`)
	rspBs := []byte(`{"original":"string","result":"string","lastName":"string","firstName":"string","middleName":"string","gender":"UNKNOWN","unparsedParts":["string"],"possible":true,"valid":true}`)
	testEndpointCall(t, reqBs, rspBs, func(client *Client) (interface{}, *http.Response, error) {
		return client.Name.Clean(context.Background(), &name.CleanRequest{
			Query: "Андрей Ильич Петров",
		})
	})
}

func TestNameSuggest(t *testing.T) {
	reqBs := []byte(`{"query":"Андре","type":"SURNAME","count":5}`)
	rspBs := []byte(`{"suggestions":[{"result":"string","lang":"string","gender":"UNKNOWN"},{"result":"string","lang":"string","gender":"UNKNOWN"},{"result":"string","lang":"string","gender":"UNKNOWN"},{"result":"string","lang":"string","gender":"UNKNOWN"},{"result":"string","lang":"string","gender":"UNKNOWN"}]}`)
	testEndpointCall(t, reqBs, rspBs, func(client *Client) (interface{}, *http.Response, error) {
		return client.Name.Suggest(context.Background(), &name.SuggestRequest{
			Type:  name.TypeSurname,
			Count: 5,
			Query: "Андре",
		})
	})
}
