package apierrors

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Error ErrorData `json:"error"`
}

type ErrorData struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func NewError(code int, description string) *Error {
	e := &Error{}
	e.Error.Code = code
	e.Error.Description = description
	return e
}

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
