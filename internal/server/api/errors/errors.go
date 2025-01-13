// Package apierrors пакет предоставляет методы работы с ошибками при работы REST API
package apierrors

import "errors"

var (
	ErrorInvalidJSONData       = errors.New("invalid JSON data")
	ErrorUserExist             = errors.New("user already exist")
	ErrorUserNotExist          = errors.New("user not exist")
	ErrorValidateLoginEmpty    = errors.New("login empty")
	ErrorValidatePasswordEmpty = errors.New("password empty")
	ErrorValidatePasswordWrong = errors.New("password wrong")
	ErrorInternalUnknown       = errors.New("unknown internal error")
	ErrorNotAuthorized         = errors.New("not authorized")
	ErrorEmptyUUID             = errors.New("empty uuid")
	ErrorNotFoundByUUID        = errors.New("not found by uuid")
)
