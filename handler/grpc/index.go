package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/devstackq/gRPC-task/pb"
	"github.com/devstackq/gRPC-task/pkg/httpclient"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedApiServiceServer
}

func RunGrpc() {
	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterApiServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

//how to move service layer ?
func (s *server) InfoByINN(ctx context.Context, input *pb.Request) (*pb.Response, error) {
	fmt.Println(input, "input arg")
	//todo: validation
	cfg := &httpclient.Config{
		BasePath: "https://suggestions.dadata.ru",
	}
	client := httpclient.New(cfg)
	resp, err := client.GetInfo(input.INN)

	if err != nil {
		return nil, err
	}
	return &pb.Response{KPP: resp.Suggestions[0].Data.KPP, FullName: resp.Suggestions[0].Data.Management.Name, Name: resp.Suggestions[0].Value, INN: input.INN}, nil
}
