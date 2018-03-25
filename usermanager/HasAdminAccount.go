package usermanager

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/risto/models"
)

func HasAdminAccount(db *gorm.DB) (bool, error) {
	var user models.User

	err := db.Where("is_admin = 1").
		First(&user).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	}

	return true, nil
}
