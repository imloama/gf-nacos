package gfnacos
import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"os"
	"path"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var nacosCfg *NacosCfg
var nacosClients *NacosClients

type NacosClients struct {
	namingClient naming_client.INamingClient
	configClient config_client.IConfigClient
}

// config file name will be: ${AppName}-${Mode}

type NacosCfg struct {
	AppName         string                      `json:"app_name"`
	Mode            string                      `json:"mode"`
	FileExtension   string                      `json:"file_extension"`
	ConfigGroup     string                      `json:"config_group"`
	DiscoveryGroup  string                      `json:"discovery_group"`
	EnableConfig    bool                        `json:"enable_config"`
	EnableDiscovery bool                        `json:"enable_discovery"`
	Meta            map[string]string           `json:"meta"`
	AppIp           string                      `json:"app_ip"`
	AppPort         uint64                      `json:"app_port"`
	Config			constant.ClientConfig 	    `json:"config"`
	Discovery 		[]constant.ServerConfig	 	`json:"discovery"`
}


// init
func Init(){
	cfg := g.Cfg()
	nacosCfg = &NacosCfg{
		//Mode: "dev",
		EnableConfig: true,
		EnableDiscovery: true,
		FileExtension: "toml",
	}
	cfg.GetStruct("nacos",nacosCfg)
	fillDefaults(nacosCfg)
	initClientConfig()

	nacosParams := vo.NacosClientParam{
		ClientConfig:  &nacosCfg.Config,
		ServerConfigs: nacosCfg.Discovery,
	}
	nacosClients = &NacosClients{}
	// 创建服务发现客户端的另一种方式
	if nacosCfg.EnableDiscovery {
		namingClient, err := clients.NewNamingClient(nacosParams)
		if err!= nil {
			fmt.Printf("create nacos naming client error! %s", err)
			panic(err)
		}
		nacosClients.namingClient = namingClient
		initDiscoveryService()
	}


	if nacosCfg.EnableConfig {
		// 创建动态配置客户端的另一种方式 (推荐)
		configClient, err := clients.NewConfigClient(nacosParams)
		if err!= nil {
			fmt.Printf("create nacos config client error! %s", err)
			panic(err)
		}
		nacosClients.configClient = configClient
		initConfigService()
	}



}

const (
	DEFAULT_DISCOVERY_GROUP = "DEFAULT_GROUP"
	DEFAULT_CONFIG_GROUP = "public"
	DEFAULT_FILE_EXTENSION = "toml"
	DEFAULT_MODE = "dev"
	DEFAULT_NACOS_CACHE_DIR = "~/nacos/cache"
	DEFAULT_NACOS
)

func fillDefaults(cfg *NacosCfg) {
	if cfg.DiscoveryGroup == "" {
		cfg.DiscoveryGroup = DEFAULT_DISCOVERY_GROUP
	}
	if cfg.ConfigGroup == "" {
		cfg.ConfigGroup = DEFAULT_CONFIG_GROUP
	}
	if cfg.FileExtension == "" {
		cfg.FileExtension = DEFAULT_FILE_EXTENSION
	}
	envMode := os.Getenv("mode")
	if envMode!="" {
		cfg.Mode = envMode
	}
	if cfg.Mode == "" {
		cfg.Mode = DEFAULT_MODE
	}
}

func initClientConfig() {
	pwd,_ := os.Getwd()
	cfg := nacosCfg.Config
	// 创建clientConfig
	cacheDir := path.Join(pwd, getPath(cfg.CacheDir, "./nacos/cache"))
	logDir   := path.Join(pwd, getPath(cfg.LogDir,"./logs/nacos"))
	cfg.CacheDir = cacheDir
	cfg.LogDir = logDir
}

func getPath(pwd,defaultPath string)string{
	if pwd == "" {
		return defaultPath
	}
	return pwd
}
