package application

type PasswordHash interface {
	Hash(password string) (string, error)
	ComparePassword(plain string, encrypted string) error
}
