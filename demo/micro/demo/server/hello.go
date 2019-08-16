package main

import (
	"context"
	"fmt"

	proto "demo/micro/demo/proto/hello"

	micro "github.com/micro/go-micro"
)

type Hello struct{}
//需要注意的是，这里的包名需要进行对应

func (h *Hello) Ping(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("helloooo"), // 服务名称
	)
	service.Init()
	proto.RegisterHelloHandler(service.Server(), new(Hello))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
