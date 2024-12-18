package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/besean163/gophkeeper/internal/client/core/api/entities"
	"github.com/besean163/gophkeeper/internal/client/tui/logger"
)

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c Client) Register(input entities.RegisterInput) (entities.TokenOutput, error) {

	var token entities.TokenOutput
	b, err := json.Marshal(input)
	if err != nil {
		logger.Get().Println(err.Error())
		return token, errors.New("ошибка зашифровки запроса")
	}
	body := bytes.NewBuffer(b)
	logger.Get().Println(body.String())

	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/register", body)
	r.Header.Add("Content-Type", "application/json")

	if err != nil {
		logger.Get().Println(err.Error())
		return token, errors.New("ошибка запроса")
	}

	response, err := http.DefaultClient.Do(r)
	if err != nil {
		logger.Get().Println(err.Error())
		return token, errors.New("ошибка ответа")
	}

	if response.StatusCode != http.StatusOK {
		return token, errors.New("ошибка сервера")
	}

	t := entities.TokenOutput{}
	rBody, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Get().Println(err.Error())
		return token, errors.New("ошибка чтения ответа")
	}

	err = json.Unmarshal(rBody, &t)
	if err != nil {
		logger.Get().Println(err.Error())
		return token, errors.New("ошибка расшифровки ответа")
	}
	return t, nil
}
