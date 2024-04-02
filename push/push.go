package push

import (
	"encoding/json"
	"fmt"
)

type Client struct {
	//app部分的client_id，这个得非常注意，这个是app区域的client_id 如果你打开华为的后台，可以看到无数个client_id，难以区分。
	ClientID string
	//app部分的client_secret
	ClientSecret string
	//项目id
	ProjectID string
}

func NewClient(clientID, clientSecret, projectID string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		ProjectID:    projectID,
	}
}

func (c Client) GetAccessToken() (map[string]any, error) {
	var data = "grant_type=client_credentials&client_id=" + c.ClientID + "&client_secret=" + c.ClientSecret
	toMap, err := Post("https://oauth-login.cloud.huawei.com/oauth2/v3/token", data, map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	})
	return toMap, err
}

func (c Client) Push(data map[string]any) (map[string]any, error) {
	token, _ := c.GetAccessToken()
	accessToken := token["access_token"].(string)
	jsonStr, _ := json.Marshal(data)
	var jsonString = string(jsonStr)
	fmt.Println(fmt.Sprintf("请求的数据：%s", jsonString))

	toMap, err := Post(fmt.Sprintf("https://push-api.cloud.huawei.com/v2/%s/messages:send", c.ProjectID), jsonString, map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	})
	return toMap, err
}
