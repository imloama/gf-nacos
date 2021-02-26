package gfnacos

import (
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var ip string
var port uint64

func initDiscoveryService(){
	ip = getServerIp()
	if nacosCfg.AppPort <= 0 {
		port = 8080
	}else{
		port = nacosCfg.AppPort
	}
	Register()
}

func getServerIp()string{
	if nacosCfg.AppIp == "" || nacosCfg.AppIp == "0.0.0.0" {
		ip,_:= getExternalIP()
		return ip
	}
	return nacosCfg.AppIp
}

func Register() (bool, error) {
	return nacosClients.namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: nacosCfg.AppName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    nacosCfg.Meta,
		//ClusterName: "cluster-a", // 默认值DEFAULT
		GroupName:  nacosCfg.DiscoveryGroup,
	})
}

func UnRegister() (bool, error) {
	return nacosClients.namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: nacosCfg.AppName,
		Ephemeral:   true,
		//Cluster:     "cluster-a", // 默认值DEFAULT
		GroupName:  nacosCfg.DiscoveryGroup,
	})
}

func GetService(name string) (model.Service, error) {
	return nacosClients.namingClient.GetService(vo.GetServiceParam{
		ServiceName: name,
		//Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		GroupName:  nacosCfg.DiscoveryGroup,
	})
}

func SelectAllInstances(name string) ([]model.Instance, error) {
	return nacosClients.namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: name,
		GroupName:  nacosCfg.DiscoveryGroup,
		//Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
	})
}

func SelectInstances(name string) ([]model.Instance, error) {
	return nacosClients.namingClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: name,
		GroupName:  nacosCfg.DiscoveryGroup,
		//Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		HealthyOnly: true,
	})
}

func SelectOneHealthyInstance(name string) (*model.Instance, error) {
	return nacosClients.namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		GroupName:  nacosCfg.DiscoveryGroup,
		//Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
	})
}