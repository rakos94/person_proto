package config

import (
	"context"
	"log"

	pb "person_proto/models"

	"google.golang.org/grpc"
)

// Ctx ...
var Ctx = context.Background()

// Client ...
var Client pb.UserServiceClient

// GrpcConn ...
var GrpcConn grpc.ClientConn

// ClientConnect ...
func ClientConnect() {
	GrpcConn, err := grpc.Dial(HostCredential, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}
	Client = pb.NewUserServiceClient(GrpcConn)
}
