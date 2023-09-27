package main

import (
	"context"
	"fmt"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Subasri-V/user-role.git/config"
	"github.com/Subasri-V/user-role.git/constants"
	"github.com/Subasri-V/user-role.git/controller"
	"github.com/Subasri-V/user-role.git/services"

	cus "github.com/Subasri-V/user-role.git/proto"
)

func initDatabase(client *mongo.Client) {
	userCollection := config.GetCollection(client, constants.DatabaseName, "User")
	roleCollection := config.GetCollection(client, constants.DatabaseName, "role")
	controller.UserRoleDetails = services.InitUserRoleService(context.Background(),userCollection,client,roleCollection)

}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)

	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	cus.RegisterUserRoleServiceServer(s,&controller.RPCServer{})
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}

