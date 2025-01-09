package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
	"github.com/besean163/gophkeeper/internal/server/api/entities/output"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

// Register запрос на регистрацию пользователя.
// Параметры:
//   - input: структура запроса.
func (c Client) Register(input input.Register) (output.Token, error) {

	var token output.Token
	b, err := json.Marshal(input)
	if err != nil {
		return token, errors.New("ошибка зашифровки запроса")
	}
	body := bytes.NewBuffer(b)

	r, err := http.NewRequest(http.MethodPost, c.Host+"/register", body)
	r.Header.Add("Content-Type", "application/json")

	if err != nil {
		return token, errors.New("ошибка запроса")
	}

	response, err := http.DefaultClient.Do(r)
	if err != nil {
		if isConnectionRefused(err) {
			return token, ErrorServerNotAvailable
		}
		return token, errors.New("ошибка ответа")
	}

	if response.StatusCode != http.StatusOK {
		var apiError apierrors.Error
		rBody, err := io.ReadAll(response.Body)
		if err != nil {
			return token, errors.New("ошибка чтения ошибки")
		}

		err = json.Unmarshal(rBody, &apiError)
		if err != nil {
			return token, errors.New("ошибка расшифровки ошибки")
		}
		return token, fmt.Errorf("ошибка сервера %d: %s", apiError.Error.Code, apiError.Error.Description)
	}

	t := output.Token{}
	rBody, err := io.ReadAll(response.Body)
	if err != nil {
		return token, errors.New("ошибка чтения ответа")
	}

	err = json.Unmarshal(rBody, &t)
	if err != nil {
		return token, errors.New("ошибка расшифровки ответа")
	}
	return t, nil
}
