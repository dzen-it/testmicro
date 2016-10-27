package main

import (
	"context"
	"log"
	"testmicro/secondservice/proto/request"

	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.secondservice"),
	)
	service.Init()

	request.RegisterSecondHandler(service.Server(), &Second{})

	if err := service.Run(); err != nil {
		log.Panic(err)
	}

}

type Second struct{}

func (s *Second) Get(ctx context.Context, req *request.GetRequest, rsp *request.GetResponse) error {
	rsp.Resp = req.Name
	return nil
}
