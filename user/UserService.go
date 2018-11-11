package user

import "github.com/google/uuid"

type UserService struct {
	currentUser User
}

func (this *UserService) CreateUser() *User {
	// if this.currentUser.GetIsAdmin() == false {
	// 	return nil
	// }

	newGuid := uuid.New()

	user := &User{
		id:      newGuid.String(),
		isAdmin: false,
	}

	return user
}
