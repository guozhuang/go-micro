package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-micro/service"
)

//从transport接收到请求
//当前层进行请求内容转化为对应srv的逻辑参数
//并将处理结果传递到transport层进行返回输出

//这样srv专注于处理对应的参数逻辑，而无需关注请求的方式是微服务还是http前端直联
//所以endpoint相当于中间层

type DiscoveryEndpoints struct {
	SayHelloEndpoint endpoint.Endpoint
	DiscoveryEndpoint endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

//sayhello 整体endpoint
type SayHelloRequest struct {
	//
}

type SayHelloResponse struct {
	Message string `json:"message"`
}

//创建招呼相关的endpoint
func MakeSayHelloEndpoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		message := srv.SayHello()
		return SayHelloResponse{Message:message}, nil
	}
}


//服务发现相关endpoint
type DiscoveryRequest struct {
	ServiceName string
}

type DiscoveryResponse struct {
	Instances []interface{} `json:"instances"`
	Error string `json:"error"`
}

func MakeDiscoveryEndpoint (srv service.Service) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (repose interface{}, err error) {
		req := request.(DiscoveryRequest)
		instances, err := srv.DiscoveryService(ctx, req.ServiceName)

		var errString = ""
		if err != nil {
			errString = err.Error()
		}

		return &DiscoveryResponse{
			Instances: instances,
			Error:     errString,
		}, nil
	}
}

//健康度检查endpoint
type HealthRequest struct {

}

type HealthRepose struct {
	Status bool `json:"status"`
}

func MakeHealthCheckEndpoint (srv service.Service) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (response interface{}, err error) {
		status := srv.HealthCheck()
		return HealthRepose{Status:status}, nil
	}
}

