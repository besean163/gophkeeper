package client

import (
	"errors"
	"testing"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/besean163/gophkeeper/internal/server/api/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	"github.com/besean163/gophkeeper/internal/server/models"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks/api/client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	httpClient := mock.NewMockHTTPClient(ctrl)
	response := mock.NewMockResponse(ctrl)
	client := client.NewClient("", httpClient, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		token     string
		mockSetup func()
		result    struct {
			item *entities.GetAccountsOutput
			err  error
		}
	}{
		{
			name: "success",
			mockSetup: func() {
				response.EXPECT().Body().Return([]byte(`{"accounts":[]}`)).Times(1)
				httpClient.EXPECT().Get("/api/accounts", map[string]string{}).Return(response, nil).Times(1)
			},
			result: struct {
				item *entities.GetAccountsOutput
				err  error
			}{
				item: &entities.GetAccountsOutput{
					Accounts: []*models.Account{},
				},
				err: nil,
			},
		},
		{
			name: "success_with_token",
			mockSetup: func() {
				response.EXPECT().Body().Return([]byte(`{"accounts":[]}`)).Times(1)
				httpClient.EXPECT().Get("/api/accounts", map[string]string{
					"Authorization": "Bearer token",
				}).Return(response, nil).Times(1)
			},
			token: "token",
			result: struct {
				item *entities.GetAccountsOutput
				err  error
			}{
				item: &entities.GetAccountsOutput{
					Accounts: []*models.Account{},
				},
				err: nil,
			},
		},
		{
			name: "fail",
			mockSetup: func() {
				httpClient.EXPECT().Get(gomock.Any(), gomock.Any()).Return(response, errors.New("test error")).Times(1)
			},
			result: struct {
				item *entities.GetAccountsOutput
				err  error
			}{
				item: nil,
				err:  errors.New("request error"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			client.SetToken(test.token)
			item, err := client.GetAccounts()
			assert.Equal(t, test.result.item, item)
			assert.Equal(t, test.result.err, err)
		})
	}

}
