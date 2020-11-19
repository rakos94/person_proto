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
	log.Println("CLIENT: GetPersonById =>", u)
	person := pb.GetPerson{
		Id: u.Id,
	}
	result, err := personService.GetPersonByID(person.Id)
	if err != nil {
		log.Println("SERVER: Failed GetPersonById =>", err)
		return &pb.Response{}, status.Errorf(codes.NotFound, "Person not found")
	}

	log.Println("SERVER: Success GetPersonById =>", u)
	data, _ := json.Marshal(&result)
	return &pb.Response{Code: "200", Data: string(data)}, nil
}
