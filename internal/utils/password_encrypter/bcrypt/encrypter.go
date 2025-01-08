package bcryptencrypter

import "golang.org/x/crypto/bcrypt"

type Encrypter struct{}

func NewEncrypter() Encrypter {
	return Encrypter{}
}

func (e Encrypter) Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (e Encrypter) CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
