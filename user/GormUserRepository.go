package user

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/risto/models"
)

type GormUserRepository struct {
	getDB func() *gorm.DB
}

func NewGormUserRepository(
	getDB func() *gorm.DB,
) GormUserRepository {
	return GormUserRepository{
		getDB: getDB,
	}
}

func (this GormUserRepository) Save(user *User) error {
	db := this.getDB()

	dbUser := models.User{
		Identification: user.GetId(),
		UserName:       user.GetUserName(),
		PasswordHash:   user.GetPasswordHash(),
		IsAdmin:        user.GetIsAdmin(),
	}

	err := db.Create(&dbUser).Error

	return err
}
