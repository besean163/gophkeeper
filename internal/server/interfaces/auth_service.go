package interfaces

type AuthService interface {
	// GetUser(login string) *models.User
	RegisterUser(login, password string) (string, error)
	LoginUser(login, password string) (string, error)
}
