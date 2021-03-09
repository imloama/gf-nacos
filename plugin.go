package gfnacos

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

const Version = "0.0.5"

type ConfigListener func(config string)

var configListener ConfigListener = func(config string) {

}


type GfNacosPlugin struct {
	Listener ConfigListener
}

func (gn GfNacosPlugin)Name()string{
	return "gf-nacos"
}

func (gn GfNacosPlugin)Author()string{
	return "github.com/imloama"
}

func (gn GfNacosPlugin)Version()string{
	return Version
}


func (gn GfNacosPlugin)Description()string{
	return "goframe and nacos"
}


func (gn GfNacosPlugin)Install(s *ghttp.Server)error{
	fmt.Println("gf-nacos插件正在安装...")
	configListener = gn.Listener
	//fmt.Printf("configListener: %s", configListener)
	return Init()
}

func (gn GfNacosPlugin)Remove()error{
	RemoveConfigListener()
	if nacosCfg.EnableDiscovery {
		UnRegister()
	}
	fmt.Println("gf-nacos插件被移除。")
	return nil
}
