package models

type TokenClaim struct {
	ID       int64
	UserType string
	Exp      int64
	Iat      int64
}
