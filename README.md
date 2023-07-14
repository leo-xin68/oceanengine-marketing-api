# oceanengine-marketing-api-go-sdk

## 概述
巨量广告 Marketing API(以下简称API) SDK 提供了Token获取、请求封装、响应解释等功能，以本地化方式轻松完成API的调用和结果的获取，旨在帮助开发者快速搭建投放管理系统。


## 如何使用
SDK数组参数调用的方法名与API接口一一对应，如
    project/list接口就对应ads.Project.List()方法
    tools/region/get接口就对应ads.Tools.Region.Get()方法
如果只想单独用某个模块，可以单独实例Tools
```go
package examples

import (
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/ads"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/config"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model"
	"github.com/leo-xin68/oceanengine-marketing-api-go-sdk/pkg/model/tools"
)

func main() {
	cfg := config.NewSdkConfig(AppId, AccessToken, false)
	oceanAds := ads.Init(cfg)

	// project/list 接口就对应 oceanAds.Project.List() 方法
	req := model.ProjectGetRequest{}
	response, headers, err := oceanAds.Project.List(*oceanAds.Ctx, req)

	// tools/region/get 接口就对应ads.Tools.Region.Get()方法
	regionReq := &tools.ToolRegionGetInfoOpts{}
	response, headers, err := oceanAds.Tools.Region.Get(*oceanAds.Ctx, regionReq)

	// 如果只想单独用某个模块，可以单独实例Tools
	oceanTools := ads.InitTools(cfg, nil)

	// tools/region/get 接口则对应 
	oceanTools.Region.Get(*oceanTools.Ctx, regionReq)

}
```


注意:model中的所有基本数据类型均为指针类型, 例如：*string, *bool, *int64, *float64