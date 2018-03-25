package main

import (
	"github.com/koodinikkarit/database"
	"github.com/koodinikkarit/risto/config"
	"github.com/koodinikkarit/risto/service"
)

func main() {
	config.ValidateConfig()

	getDB := database.CreateGetDB(
		config.MysqlUsername,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDatabase,
	)

	service.StartRistoService(getDB, config.ServicePort)

	// hash, _ := authentication.CreatePasswordHash("kameli38")

	//fmt.Println("hash is", hash)

	// payload := `{"hello": "world"}`

	// token, err := jose.Sign(payload, jose.NONE, nil)

	// jose.SignBytes()

	// if err == nil {
	// 	fmt.Printf("\nHS256 = %v\n", token)
	// 	s, pay, err := jose.Decode(token, nil)
	// 	fmt.Println("Decoded payload", pay, s, err)
	// } else {
	// 	fmt.Println(err)
	// }
}
