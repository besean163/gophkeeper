package client

import "errors"

var (
	ErrorServerNotAvailable = errors.New("server not available")
	ErrorRequestError       = errors.New("request error")
	ErrorServerError        = errors.New("server error")
)
