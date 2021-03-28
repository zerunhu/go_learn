package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	api_url := "https://oapi.dingtalk.com/robot/send?access_token=66a48ed58142ef88953fc67b794bc1b5d720f8e1a2a78fb6637c6485ce14dd5a"
	content := `{"msgtype": "text",
      "text": {"content": "123rrr"},
                "at": {
                     "atMobiles": [
                         "18301371817"
                     ],
                     "isAtAll": false
                }
    }`
	//dataType, _ := json.Marshal(b)
	////var jsonstr = []byte(dataType)
	//buffer := bytes.NewBuffer(dataType)
	request, err := http.NewRequest("POST", api_url, strings.NewReader(content))
	if err != nil {
		fmt.Printf("http.NewRequest%v", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	client := http.Client{}                                              //创建客户端
	resp, err := client.Do(request.WithContext(context.TODO()))          //发送请求
	if err != nil {
		fmt.Printf("client.Do%v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll%v", err)
	}
	fmt.Println(string(respBytes))
}
