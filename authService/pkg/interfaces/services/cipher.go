package services

type ICipherService interface {
	Decrypt(message string) (decMessage string, err error)
	Encrypt(message string) (encMessage string, err error)
}
