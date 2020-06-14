package service

import (
	"context"
	"errors"
	"go-micro/config"
	"go-micro/discover"
)

type Service interface {
	//健康检查
	HealthCheck() bool

	//
	SayHello() string

	//服务发现接口
	DiscoveryService(ctx context.Context, serviceName string)([]interface{}, error)
}

//当前包内的标准错误
var ErrNotServiceInstances = errors.New("instances are not existed")

func NewDiscoveryServiceImpl(discoveryClient discover.DiscoveryClient) Service  {
	return &DiscoveryServiceImpl{
		discoveryClient:discoveryClient,
	}
}

//接口的具体实现
type DiscoveryServiceImpl struct {
	discoveryClient discover.DiscoveryClient
}

func (*DiscoveryServiceImpl) SayHello() string {
	return "Hello World!"
}

func (service *DiscoveryServiceImpl) DiscoveryService (ctx context.Context, serviceName string) ([]interface{}, error) {

	//从服务中获取列表
	instances := service.discoveryClient.DiscoverServices(serviceName, config.Logger)

	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}

	return instances, nil
}

func (*DiscoveryServiceImpl) HealthCheck() bool {
	return true
}