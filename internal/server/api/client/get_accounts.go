package client

import (
	"encoding/json"
	"errors"

	"github.com/besean163/gophkeeper/internal/server/api/entities/output"
)

// GetAccounts запрос на получение списка аккаунтов.
func (c Client) GetAccounts() (*output.GetAccounts, error) {
	var output *output.GetAccounts
	response, err := c.Get(c.Host + "/api/accounts")

	if err != nil {
		return nil, errors.New("request error")
	}

	err = json.Unmarshal(response.Body(), &output)
	if err != nil {
		return nil, errors.New("read answer error")
	}

	return output, nil
}
