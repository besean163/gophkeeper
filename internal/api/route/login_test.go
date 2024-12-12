package route

// func TestLoginRoute(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	authService := mock.NewMockAuthService(ctrl)
// 	authService.EXPECT().GetUser("test_user")

// 	tests := []struct {
// 		name       string
// 		method     string
// 		body       string
// 		exceptCode int
// 	}{
// 		{
// 			name:       "success_create",
// 			method:     http.MethodPost,
// 			body:       `{"login":"create_user","password":"create_password"}`,
// 			exceptCode: http.StatusOK,
// 		},
// 		{
// 			name: "test_2",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			assert.Equal(t, test.name, test.name)
// 		})
// 	}

// }
