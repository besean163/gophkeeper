// Package client клиентская библиотека для работы с сервером
package client

import (
	"errors"
	"net"

	"github.com/besean163/gophkeeper/internal/logger"
	httpclient "github.com/besean163/gophkeeper/internal/server/api/client/http_client"
)

// Client входная структура для работы.
type Client struct {
	httpclient.HTTPClient
	logger logger.Logger
	// Token токен клиента
	Token string
	// Host базовый хост для запросов
	Host string
}

// NewClient создать клиент
func NewClient(host string, httpClient httpclient.HTTPClient, logger logger.Logger) *Client {
	return &Client{
		Host:       host,
		logger:     logger,
		HTTPClient: httpClient,
	}
}

// SetToken установить токен для запросов
func (c *Client) SetToken(token string) {
	c.Token = token
}

func isConnectionRefused(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return opErr.Err.Error() == "connect: connection refused"
	}
	return false
}

// Get прямой "Get" запрос к серверу.
// Параметры:
//   - uri: uri запроса.
func (c Client) Get(uri string) (httpclient.Response, error) {
	headers := make(map[string]string, 0)
	headers = c.setTokenHeader(headers)

	return c.HTTPClient.Get(uri, headers)
}

// Post прямой "Post" запрос к серверу.
// Параметры:
//   - uri: uri запроса.
//   - body: тело запроса.
func (c Client) Post(uri string, body interface{}) (httpclient.Response, error) {
	headers := make(map[string]string, 0)
	headers = c.setTokenHeader(headers)
	headers = c.setJsonHeader(headers)

	return c.HTTPClient.Post(uri, body, headers)
}

// Put прямой "Put" запрос к серверу.
// Параметры:
//   - uri: uri запроса.
//   - body: тело запроса.
func (c Client) Put(uri string, body interface{}) (httpclient.Response, error) {
	headers := make(map[string]string, 0)
	headers = c.setTokenHeader(headers)
	headers = c.setJsonHeader(headers)

	return c.HTTPClient.Put(uri, body, headers)
}

// Delete прямой "Delete" запрос к серверу.
// Параметры:
//   - uri: uri запроса.
//   - body: тело запроса.
func (c Client) Delete(uri string, body interface{}) (httpclient.Response, error) {
	headers := make(map[string]string, 0)
	headers = c.setTokenHeader(headers)
	headers = c.setJsonHeader(headers)

	return c.HTTPClient.Delete(uri, body, headers)
}

func (c Client) setTokenHeader(headers map[string]string) map[string]string {
	if c.Token != "" {
		headers["Authorization"] = "Bearer " + c.Token
	}
	return headers
}

func (c Client) setJsonHeader(headers map[string]string) map[string]string {
	headers["Content-Type"] = "application/json"
	return headers
}
