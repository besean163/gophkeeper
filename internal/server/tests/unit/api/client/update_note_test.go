package client

import (
	"errors"
	"net/http"
	"testing"

	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/besean163/gophkeeper/internal/server/api/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	mock "github.com/besean163/gophkeeper/internal/server/tests/mocks/api/client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	httpClient := mock.NewMockHTTPClient(ctrl)
	response := mock.NewMockResponse(ctrl)
	c := client.NewClient("", httpClient, defaultlogger.NewDefaultLogger())

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
				httpClient.EXPECT().Put("/api/note", entities.NoteUpdateInput{}, map[string]string{
					"Content-Type": "application/json",
				}).Return(response, nil).Times(1)
			},
		},
		{
			name: "success_with_token",
			mockSetup: func() {
				response.EXPECT().StatusCode().Return(http.StatusOK).Times(1)
				httpClient.EXPECT().Put("/api/note", entities.NoteUpdateInput{}, map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "Bearer token",
				}).Return(response, nil).Times(1)
			},
			token: "token",
		},
		{
			name: "fail",
			mockSetup: func() {
				httpClient.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()).Return(response, errors.New("test error")).Times(1)
			},
			result: client.ErrorRequestError,
		},
		{
			name: "wrong_status_code",
			mockSetup: func() {
				response.EXPECT().StatusCode().Return(http.StatusBadRequest).Times(1)
				httpClient.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()).Return(response, nil).Times(1)
			},
			result: client.ErrorServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockSetup()
			c.SetToken(test.token)
			err := c.UpdateNote(entities.NoteUpdateInput{})
			assert.Equal(t, test.result, err)
		})
	}

}
