package service

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/risto/authentication"
	"golang.org/x/net/context"

	"github.com/koodinikkarit/go-clientlibs/risto"
	"github.com/koodinikkarit/risto/generators"
	"github.com/koodinikkarit/risto/models"
)

func (s *RistoServiceServer) CreateToken(
	ctx context.Context,
	in *RistoService.CreateTokenRequest,
) (
	*RistoService.CreateTokenResponse,
	error,
) {
	res := &RistoService.CreateTokenResponse{}

	db := s.getDB()

	var user models.User

	err := db.Find(&user, "user_name = ?", in.Username).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return res, err
	}

	if user.ID == 0 {
		res.State = RistoService.CreateTokenResponse_USER_NOT_FOUND
		return res, nil
	}

	passwordCorrect := authentication.ComparePsswordHashToPlainPassword(user.PasswordHash, in.Password)

	if passwordCorrect == false {
		res.State = RistoService.CreateTokenResponse_PASSWORD_WRONG
		return res, nil
	}

	token, err := authentication.CreateToken(user)

	if err != nil {
		return res, err
	}

	res.Token = token
	res.State = RistoService.CreateTokenResponse_TOKEN_CREATED
	res.User = generators.NewProtoUser(user)

	return res, nil
}

func (s *RistoServiceServer) VerifyToken(
	ctx context.Context,
	in *RistoService.VerifyTokenRequest,
) (
	*RistoService.VerifyTokenResponse,
	error,
) {
	res := &RistoService.VerifyTokenResponse{}

	return res, nil
}

func (s *RistoServiceServer) FetchUserByToken(
	ctx context.Context,
	in *RistoService.FetchUserByTokenRequest,
) (
	*RistoService.FetchUserByTokenResponse,
	error,
) {
	res := &RistoService.FetchUserByTokenResponse{}

	return res, nil
}
