package pencrypt

type Encrypter interface {
	Encrypt(password string) (string, error)
	CheckPassword(hash, password string) bool
}
