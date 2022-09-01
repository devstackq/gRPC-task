package service

import (
	"context"
	"fmt"

	"github.com/devstackq/gRPC-task/pb"
	"github.com/devstackq/gRPC-task/pkg/httpclient"
)

type Service struct {
	pb.UnimplementedApiServiceServer
}
type ServiceInterface interface {
	InfoByINN(ctx context.Context, input *pb.Request) (*pb.Response, error)
}

func (s *Service) InfoByINN(ctx context.Context, input *pb.Request) (*pb.Response, error) {
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
