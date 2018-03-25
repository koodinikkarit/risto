package generators

import (
	"github.com/koodinikkarit/go-clientlibs/risto"
	"github.com/koodinikkarit/risto/models"
)

func NewProtoUser(user models.User) *RistoService.User {
	return &RistoService.User{
		Id:       user.ID,
		UserName: user.UserName,
		IsAdmin:  user.IsAdmin,
	}
}
