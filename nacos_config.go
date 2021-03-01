package gfnacos

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcfg"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func initConfigService(){
	if !nacosCfg.EnableConfig {
		return
	}
	dataId := fmt.Sprintf("%s-%s.%s", nacosCfg.AppName, nacosCfg.Mode, nacosCfg.FileExtension)
	content, err := nacosClients.configClient.GetConfig(vo.ConfigParam{
		Group: nacosCfg.ConfigGroup,
		DataId: dataId,
	})
	if err!=nil{
		g.Log().Errorf("读取nacos配置文件[Group=%s,dataId=%s]发生错误！%s", nacosCfg.ConfigGroup, dataId, err)
	}
	setGcfgContent(content)

	nacosClients.configClient.ListenConfig(vo.ConfigParam{
		Group: nacosCfg.ConfigGroup,
		DataId: dataId,
		OnChange: onConfigChange,
	})

}

func onConfigChange(namespace, group, dataId, data string) {
	g.Log().Printf("nacos config change, namespace=%s, group=%s, dataId=%s, data=%s", namespace, group, dataId, data)
	setGcfgContent(data)
}

func setGcfgContent(content string){
	g.Log().Debugf("got config file content: %s", content)
	if content!=""{
		g.Cfg().Clear()
		gcfg.SetContent(content)
	}
}

func RemoveConfigListener(){
	if !nacosCfg.EnableConfig {
		return
	}
	dataId := fmt.Sprintf("%s-%s.%s", nacosCfg.AppName, nacosCfg.Mode, nacosCfg.FileExtension)
	err := nacosClients.configClient.CancelListenConfig(vo.ConfigParam{
		Group: nacosCfg.ConfigGroup,
		DataId: dataId,
	})
	if err!=nil{
		g.Log().Errorf("取消nacos配置监听失败！%s", err)
	}
}