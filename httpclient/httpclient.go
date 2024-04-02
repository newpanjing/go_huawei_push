package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func StreamToMap(reader io.ReadCloser) (map[string]any, error) {
	defer reader.Close()
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	str := string(body)
	//转为json
	var dataMap = make(map[string]any)
	err = json.Unmarshal([]byte(str), &dataMap)
	if err != nil {
		return nil, err
	}

	return dataMap, nil
}

func Do(req *http.Request) (map[string]any, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return StreamToMap(resp.Body)
}
func Post(url string, data string, header map[string]string) (map[string]any, error) {

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	for key := range header {
		req.Header.Set(key, header[key])
	}
	return Do(req)
}
