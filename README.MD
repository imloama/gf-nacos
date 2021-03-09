# goframe nacos plugin

使用nacos作为配置中心和服务注册管理

`version:0.0.6` 

默认使用toml配置文件

假设`config.toml`配置
```
[nacos]
    AppName = "example"
    Mode= "dev"
    FileExtension="toml"
```
那么会去nacos中查询`${AppName}-${Mode}.${FileExtension}`即`example-dev.toml`的配置文件，
读取到配置内容后会覆盖g.Cfg中的内容，调用方法：`gcfg.SetContent(content)`，

注意：

***更新前会调用clear()，清理掉内存中之前读取的配置文件***




## 示例 

[config.toml](example/config.toml)

[main.go](example/main.go)


```toml

[nacos]
    AppName = "example"
    Mode    = "dev"
    FileExtension = "toml"
    EnableConfig = true
    EnableDiscovery = true
    AppPort = 9000

[nacos.config]
    TimeoutMs = 5000
    NotLoadCacheAtStart = true
    RotateTime = "1h"
    MaxAge = 3
    LogLevel = "debug"

[nacos.discovery]
    Scheme = "http"
    ContextPath = "/nacos"
    IpAddr = "127.0.0.1"
    Port  = 8848

```

## 引用示例
```go

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/imloama/gf-nacos"
)

func main(){
	g.Cfg().SetPath("./")
    s := g.Server()
    s.BindHandler("/", func(r *ghttp.Request) {
        r.Response.Write("hello, nacos!")
    })
    s.Plugin(&gfnacos.GfNacosPlugin{
        Listener: func(config string) {
            g.Log().Println("配置文件发生了更新！")
        },
    })
    s.Run()
}


```