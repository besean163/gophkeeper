package client

import (
	"errors"
	"net/http"
	"testing"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/besean163/gophkeeper/internal/server/api/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks/api/client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	httpClient := mock.NewMockHTTPClient(ctrl)
	response := mock.NewMockResponse(ctrl)
	client := client.NewClient("", httpClient, defaultlogger.NewDefaultLogger())

	tests := []struct {
		name      string
		token     string
		mockSetup func()
		result    error
	}{
		{
			name: "success",
			mockSetup: func() {
				response.EXPECT().StatusCode().Return(http.StatusOK).Times(1)
				httpClient.EXPECT().Post("/api/card", input.CardCreate{}, map[string]string{
					"Content-Type": "application/json",
				}).Return(response, nil).Times(1)
			},
		},
		{
			name: "success_with_token",
			mockSetup: func() {
				response.EXPECT().StatusCode().Return(http.StatusOK).Times(1)
				httpClient.EXPECT().Post("/api/card", input.CardCreate{}, map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "Bearer token",
				}).Return(response, nil).Times(1)
			},
			token: "token",
		},
		{
			name: "fail",
			mockSetup: func() {
				httpClient.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(response, errors.New("test error")).Times(1)
			},
			result: errors.New("request error"),
		},
		{
			name: "wrong_status_code",
			mockSetup: func() {
				response.EXPECT().StatusCode().Return(http.StatusBadRequest).Times(1)
				httpClient.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(response, nil).Times(1)
			},
			result: errors.New("server error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			client.SetToken(test.token)
			err := client.CreateCard(input.CardCreate{})
			assert.Equal(t, test.result, err)
		})
	}

}
