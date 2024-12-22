package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/besean163/gophkeeper/internal/client/core/api/entities"
	"github.com/besean163/gophkeeper/internal/logger"
)

const serverURL = "http://localhost:8080"

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c Client) Register(input entities.GetTokenInput) (entities.TokenOutput, error) {

	var token entities.TokenOutput
	b, err := json.Marshal(input)
	if err != nil {
		logger.Get().Println(err.Error())
		return token, errors.New("ошибка зашифровки запроса")
	}
	body := bytes.NewBuffer(b)
	logger.Get().Println(body.String())

	r, err := http.NewRequest(http.MethodPost, serverURL+"/register", body)
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

func (c Client) Login(input entities.GetTokenInput) (entities.TokenOutput, error) {
	var token entities.TokenOutput
	b, err := json.Marshal(input)
	if err != nil {
		logger.Get().Println(err.Error())
		return token, errors.New("ошибка зашифровки запроса")
	}
	body := bytes.NewBuffer(b)
	logger.Get().Println(body.String())

	r, err := http.NewRequest(http.MethodPost, serverURL+"/login", body)
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

func (c Client) GetAccounts(token string) (entities.AccountsOutput, error) {
	output := entities.AccountsOutput{}

	r, err := http.NewRequest(http.MethodGet, serverURL+"/api/accounts", nil)
	r.Header.Add("Authorization", "Bearer "+token)

	if err != nil {
		logger.Get().Println(err.Error())
		return output, errors.New("ошибка запроса")
	}

	response, err := http.DefaultClient.Do(r)
	if err != nil {
		logger.Get().Println(err.Error())
		return output, errors.New("ошибка ответа")
	}

	if response.StatusCode != http.StatusOK {
		return output, errors.New("ошибка сервера")
	}

	rBody, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Get().Println(err.Error())
		return output, errors.New("ошибка чтения ответа")
	}

	err = json.Unmarshal(rBody, &output)
	if err != nil {
		logger.Get().Println(err.Error())
		return output, errors.New("ошибка расшифровки ответа")
	}

	return output, nil
}
