package usermanager

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/risto/authentication"

	"github.com/koodinikkarit/risto/models"
)

func CreateAdmin(
	db *gorm.DB,
	password string,
) (
	models.User,
	error,
) {
	user := models.User{
		IsAdmin:  true,
		UserName: "admin",
	}

	passwordHash, err := authentication.CreatePasswordHash(password)

	if err != nil {
		return user, err
	}

	user.PasswordHash = passwordHash

	err = db.Create(&user).Error

	if err != nil {
		return user, nil
	}

	return user, nil
}
