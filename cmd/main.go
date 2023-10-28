package main

import (
	"fmt"
	"log"
	"net"
	"server/pkg/common"
	"server/pkg/pb"
	"server/pkg/repository"
	"server/pkg/server"

	"google.golang.org/grpc"
)

func main() {
	config, err := common.LoadConfig()
	if err != nil {
		log.Println("error loading config :", err)
		return
	}

	services := IniApi(config)

	server, listner := InitServer(config, services)

	fmt.Println("server running on port", config.Port)
	server.Serve(listner)

}

func IniApi(cfg common.Config) pb.UserServer {
	repo := repository.NewRepository()
	services := server.NewUserServices(repo)
	return services
}

func InitServer(cfg common.Config, services pb.UserServer) (*grpc.Server, net.Listener) {
	server := grpc.NewServer()
	listner, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	pb.RegisterUserServer(server, services)
	return server, listner
}
