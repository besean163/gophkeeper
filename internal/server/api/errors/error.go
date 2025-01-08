package apierrors

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error структура ошибки в ответе REST API.
type Error struct {
	Error ErrorData `json:"error"`
}

// ErrorData детальные данные ошибки.
type ErrorData struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// NewError создание ошибки.
func NewError(code int, description string) *Error {
	e := &Error{}
	e.Error.Code = code
	e.Error.Description = description
	return e
}

// SendError запись ошибки в тело ответа.
func SendError(w http.ResponseWriter, code int, description string) {
	apiError := Error{}
	apiError.Error.Code = code
	apiError.Error.Description = description

	errorBody, err := json.Marshal(apiError)
	if err != nil {
		log.Println("send error fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(apiError.Error.Code)
	w.Write(errorBody)
}
