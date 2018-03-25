package models

type User struct {
	ID           uint64
	UserName     string
	PasswordHash []byte
	IsAdmin      bool
}
