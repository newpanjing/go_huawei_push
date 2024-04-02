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
	m, err := client.Push(dataMap)
	if err != nil {
		return
	}
	fmt.Println(m)
}
