package authentication

import (
	"encoding/json"

	"github.com/dvsekhvalnov/jose2go"
	"github.com/koodinikkarit/risto/models"
)

type UserToken struct {
	UserName string
}

var jwtKey = []byte{97, 48, 97, 50, 97, 98, 100, 56, 45, 54, 49, 54, 50, 45, 52, 49, 99, 51, 45, 56, 51, 100, 54, 45, 49, 99, 102, 53, 53, 57, 98, 52, 54, 97, 102, 99}

func CreateToken(user models.User) (string, error) {
	userToken := UserToken{
		UserName: user.UserName,
	}

	tokenBytes, err := json.Marshal(userToken)

	if err != nil {
		return "", err
	}

	return jose.SignBytes(tokenBytes, jose.HS256, jwtKey)
}

func DecodeUser(token string) models.User {
	payload, _, _ := jose.Decode(token, jwtKey)

	var userToken UserToken

	json.Unmarshal([]byte(payload), &userToken)

	return models.User{
		UserName: userToken.UserName,
	}
}
