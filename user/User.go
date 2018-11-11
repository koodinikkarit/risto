package user

import "github.com/koodinikkarit/risto/authentication"

type User struct {
	id           string
	userName     string
	passwordHash []byte
	isAdmin      bool
}

func (this *User) GetId() string {
	return this.id
}

func (this *User) GetIsAdmin() bool {
	return this.isAdmin
}

func (this *User) ChangeUserToAdmin() {
	this.isAdmin = true
}

func (this *User) ChangeUserToNormalUser() {
	this.isAdmin = false
}

func (this *User) ChangePassword(
	newPassowrd string,
) error {
	passwordHash, err := authentication.CreatePasswordHash(newPassowrd)

	if err == nil {
		this.passwordHash = passwordHash
	}

	return err
}

func (this *User) GetPasswordHash() []byte {
	return this.passwordHash
}

func (this *User) GetUserName() string {
	return this.userName
}

func (this *User) ChangeUserName(newUserName string) {
	this.userName = newUserName
}
