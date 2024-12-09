package route

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock "github.com/besean163/gophkeeper/internal/auth/mocks"
	"github.com/besean163/gophkeeper/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoute(t *testing.T) {

	ctrl := gomock.NewController(t)
	authService := mock.NewMockAuthService(ctrl)
	authService.EXPECT().GetUser("create_user").Return(nil).Times(1)
	authService.EXPECT().GetUser("exist_user").Return(&models.User{
		Login: "exist_user",
	}).Times(1)

	handler := RegisterRoute(authService)
	tests := []struct {
		name           string
		method         string
		inputBody      string
		outputBody     string
		exceptCode     int
		requestHeaders map[string]string
	}{
		{
			name:       "fail_without_json_header",
			method:     http.MethodPost,
			inputBody:  `{"login":"create_user","password":"create_password"}`,
			outputBody: `expect JSON data`,
			exceptCode: http.StatusBadRequest,
		},
		{
			name:       "success_create",
			method:     http.MethodPost,
			inputBody:  `{"login":"create_user","password":"create_password"}`,
			exceptCode: http.StatusOK,
			requestHeaders: map[string]string{
				"Content-type": "application/json",
			},
		},
		{
			name:       "fail_create_user_exist",
			method:     http.MethodPost,
			inputBody:  `{"login":"exist_user","password":"exist_password"}`,
			outputBody: `user already exist`,
			exceptCode: http.StatusBadRequest,
			requestHeaders: map[string]string{
				"Content-type": "application/json",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inputBody := strings.NewReader(test.inputBody)
			request, _ := http.NewRequest(test.method, "/api/user/register", inputBody)
			if test.requestHeaders != nil {
				for h, v := range test.requestHeaders {
					request.Header.Set(h, v)
				}
			}
			// request.Header.Set("Content-type", "application/json")
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, request)

			response := rr.Result()
			defer response.Body.Close()
			assert.Equal(t, test.exceptCode, response.StatusCode)

			outputBody, _ := io.ReadAll(response.Body)
			assert.Equal(t, test.outputBody, string(outputBody))
			// fmt.Println(string(outputBody))
		})
	}

}
