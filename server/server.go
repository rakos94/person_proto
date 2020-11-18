package server

import (
	"context"
	"encoding/json"
	"log"
	"person_proto/pb"
	"person_proto/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var personService services.PersonService = services.PersonServiceImpl{}

type Server struct {
	pb.UnimplementedPersonServiceServer
}

func (*Server) GetPersonById(ctx context.Context, u *pb.GetPerson) (*pb.Response, error) {
	log.Println("Getperson from client =>", u)
	person := pb.GetPerson{
		Id: u.Id,
	}
	result, err := personService.GetPersonByID(person.Id)
	if err != nil {
		log.Println("Failed get person err =>", err)
		return &pb.Response{}, status.Errorf(codes.NotFound, "Person not found")
	}

	log.Println("Success get person from client =>", u)
	data, _ := json.Marshal(&result)
	return &pb.Response{Code: "200", Data: string(data)}, nil
}
