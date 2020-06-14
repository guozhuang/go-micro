package discover

import "log"

//服务发现客户端
type DiscoveryClient interface {
	//服务注册
	Register (serviceName, instanceId, healthCheckUrl string,
		instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool
	//服务注销
	DeRegister (instanceId string, logger *log.Logger) bool
	//服务发现
	DiscoverServices (serviceName string, logger *log.Logger) [] interface{}
}
