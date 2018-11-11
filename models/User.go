package models

type User struct {
	ID             uint64
	Identification string
	UserName       string
	PasswordHash   []byte
	IsAdmin        bool
}
