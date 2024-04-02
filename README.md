# 华为推送golang服务端实现

实现这个的包的目的是降低开发难度，而且官方的也不是一个标准的golang包，没法直接引用。

所以一不做二不休就自己实现了一个基于api直接调用的包，代码复杂度低很多。

而且官方实现的包写得和屎💩一样



## 流程

首先需要去华为的AppGallery注册账号，然后创建应用，获取appId和appSecret，然后就可以调用接口了。

然后需要再app上进行jks加密，然后上传到华为的appGallery，sha256加密的key

最后启动app 获取token就ok了。

```golang
package example

import (
	"fmt"
	"github.com/newpanjing/go_huawei_push/push"
	"testing"
	"time"
)

func Test(t *testing.T) {
	//替换为自己的appId和appSecret
	const appClientID = "110****81"
	const appClientSecret = "7b47d3*******655"
	//替换为自己的projectID
	const projectID = "38*****8"
	client := push.NewClient(appClientID, appClientSecret, projectID)

	var dataMap = map[string]any{
		"validate_only": false,
		"message": map[string]any{
			"android": map[string]any{
				//每日500条，没有上架的应用
				//"target_user_type": 1,
				//https://huaweicloud.csdn.net/643f650b7de2bc0e53e30881.html
				//https://huaweicloud.csdn.net/643f650b7de2bc0e53e30881.html
				//申请自分类后不受限制
				"category": "IM",
				"notification": map[string]any{
					"bi_tag": "test title12",
					"tag":    time.Now().String(),
					"title":  "test title12",
					"body":   fmt.Sprintf("当前时间为：%s", time.Now().Format("2006-01-02 15:04:05")),
					"ttl":    "1296000s",
					//"urgency":  "HIGH",
					"click_action": map[string]any{
						"type": 3,
					},
				},
			},
			"token": []string{
				"从手机中获取的token",
			},
		},
	}
	//dataMap 中请完全参考官网文档：
	//https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/rest-sample-code-0000001050040242
	//
	m, err := client.Push(dataMap)
	if err != nil {
		return
	}
	fmt.Println(m)
}

```