// Package httpclient предоставляет кастомные интерфейсы клиента http запросов и ответа
package httpclient

// Response кастомный интерфейс ответа.
type Response interface {
	StatusCode() int
	Body() []byte
}

// HTTPClient кастомный интерфейс клиента.
type HTTPClient interface {
	Get(uri string, headers map[string]string) (Response, error)
	Post(uri string, body interface{}, headers map[string]string) (Response, error)
	Put(uri string, body interface{}, headers map[string]string) (Response, error)
	Delete(uri string, body interface{}, headers map[string]string) (Response, error)
}
