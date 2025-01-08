// Package restyclient http клиент основанный на библиотеке resty
package restyclient

import (
	httpclient "github.com/besean163/gophkeeper/internal/server/api/client/http_client"
	"github.com/go-resty/resty/v2"
)

type HTTPClient struct {
	resty.Client
}

func NewHTTPClient() HTTPClient {
	return HTTPClient{
		Client: *resty.New(),
	}
}

func (c HTTPClient) Get(uri string, headers map[string]string) (httpclient.Response, error) {
	return c.Client.R().SetHeaders(headers).Get(uri)
}

func (c HTTPClient) Post(uri string, body interface{}, headers map[string]string) (httpclient.Response, error) {
	return c.Client.R().SetHeaders(headers).SetBody(body).Post(uri)
}

func (c HTTPClient) Put(uri string, body interface{}, headers map[string]string) (httpclient.Response, error) {
	return c.Client.R().SetHeaders(headers).SetBody(body).Put(uri)
}

func (c HTTPClient) Delete(uri string, body interface{}, headers map[string]string) (httpclient.Response, error) {
	return c.Client.R().SetHeaders(headers).SetBody(body).Delete(uri)
}
