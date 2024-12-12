package apierrors

import "errors"

var (
	ErrorInvalidJSONData       = errors.New("invalid JSON data")
	ErrorRegisterUserExist     = errors.New("user already exist")
	ErrorLoginUserNotExist     = errors.New("user not exist")
	ErrorValidateLoginEmpty    = errors.New("login empty")
	ErrorValidatePasswordEmpty = errors.New("password empty")
	ErrorInternalUnknown       = errors.New("unknown internal error")
)
