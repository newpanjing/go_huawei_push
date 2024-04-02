package main

import (
	"encoding/json"
	"fmt"
	"huawei_push/httpclient"
	"time"
)

// 应用部分的，不是项目部分的，很坑

const appClientID = "110678581"
const appClientSecret = "7b47d3eecbe0beb19600649fcba51b96568a61cb43424f9fac289ea02c10c655"
const projectID = "388421841222062598"

func getAccessToken() (map[string]any, error) {

	var data = "grant_type=client_credentials&client_id=" + appClientID + "&client_secret=" + appClientSecret
	toMap, err := httpclient.Post("https://oauth-login.cloud.huawei.com/oauth2/v3/token", data, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	})
	return toMap, err
}

func Push() (map[string]any, error) {
	token, _ := getAccessToken()
	accessToken := token["access_token"].(string)
	fmt.Println(accessToken)

	var jsonMap = map[string]any{
		"validate_only": false,
		"message": map[string]any{
			"notification": map[string]any{
				"category": "IM",
				"title":    "test title11",
				"body":     "test body11",
			},
			"android": map[string]any{
				//每日500条，没有上架的应用
				//https://huaweicloud.csdn.net/643f650b7de2bc0e53e30881.html
				//
				"target_user_type": 1,
				"notification": map[string]any{
					"category": "IM",
					"bi_tag":   "test title12",
					"tag":      time.Now().String(),
					"title":    "test title12",
					"body":     fmt.Sprintf("当前时间为：%s", time.Now().Format("2006-01-02 15:04:05")),
					"ttl":      "1296000s",
					//"urgency":  "HIGH",
					"click_action": map[string]any{
						"type": 3,
					},
				},
			},
			"token": []string{
				"IQAAAACy05-lAAD22kuFdChDlBZom18l134kt2ehgA6h3KVzOg9efnb_f71ivUQsXd54L_fHRSvDp8Xj3kZDQz-03DaQB13y-G8i7Gs6EzqL4k2a7g",
			},
		},
	}

	//json to string
	data, _ := json.Marshal(jsonMap)
	var jsonString = string(data)
	fmt.Println(fmt.Sprintf("请求的数据：%s", jsonString))
	toMap, err := httpclient.Post(fmt.Sprintf("https://push-api.cloud.huawei.com/v2/%s/messages:send", projectID), jsonString, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	})
	return toMap, err

}

func main() {

	for i := 0; i < 10; i++ {

		push, err := Push()
		if err != nil {
			return
		}
		fmt.Println(push)

		time.Sleep(time.Second * 1)
	}
}
