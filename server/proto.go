package server

import (
	"log"
	"net"
	"person_proto/config"
	"person_proto/pb"

	grpc "google.golang.org/grpc"
)

// SetProto for init proto server
func SetProto() {
	lis, err := net.Listen("tcp", ":"+config.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen with err =>", err)
	}

	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve with err =>", err)
	}
}
