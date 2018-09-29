package service

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/go-clientlibs/risto"
	"github.com/koodinikkarit/risto/generators"
	"github.com/koodinikkarit/risto/usermanager"
)

func (s *RistoServiceServer) CreateUser(
	ctx context.Context,
	in *RistoService.CreateUserRequest,
) (
	*RistoService.CreateUserResponse,
	error,
) {
	res := &RistoService.CreateUserResponse{}

	return res, nil
}

func (s *RistoServiceServer) HasAdminAccount(
	ctx context.Context,
	in *RistoService.HasAdminAccountRequest,
) (
	*RistoService.HasAdminAccountResponse,
	error,
) {
	res := &RistoService.HasAdminAccountResponse{}

	fmt.Println("HasAdminAccount")

	hasAdminAccount, err := usermanager.HasAdminAccount(s.getDB())

	res.HasAdminAccount = hasAdminAccount

	return res, err
}

func (s *RistoServiceServer) CreateAdminAccount(
	ctx context.Context,
	in *RistoService.CreateAdminAccountRequest,
) (
	*RistoService.CreateAdminAccountResponse,
	error,
) {
	res := &RistoService.CreateAdminAccountResponse{}

	db := s.getDB()

	hasAdminAccount, err := usermanager.HasAdminAccount(db)

	fmt.Println("err", err)

	if err != nil && err != gorm.ErrRecordNotFound {
		res.Success = false
		return res, err
	}

	if hasAdminAccount == true {
		res.Success = false
		return res, nil
	}

	user, err := usermanager.CreateAdmin(
		s.getDB(),
		in.Password,
	)

	res.User = generators.NewProtoUser(user)
	res.Success = err == nil

	return res, err
}
