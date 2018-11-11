package service

import (
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/go-clientlibs/risto"
	"github.com/koodinikkarit/risto/user"
)

type RistoServiceServer struct {
	getDB          func() *gorm.DB
	userService    *user.UserService
	userRepository user.UserRepository
}

func StartRistoService(
	getDB func() *gorm.DB,
	port string,
) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	userRepository := user.NewGormUserRepository(getDB)

	RistoService.RegisterRistoServer(s, &RistoServiceServer{
		getDB:          getDB,
		userRepository: userRepository,
	})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
