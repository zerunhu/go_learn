package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func ListSplitStatus1(alerts []interface{}) map[string][]interface{} {
	var firing []interface{}
	var resolved []interface{}
	for _, alert := range alerts {
		if alert.(map[string]interface{})["status"] == "firing" {
			firing = append(firing, alert)
		} else {
			resolved = append(resolved, alert)
		}
	}
	SplitStatus := make(map[string][]interface{})
	SplitStatus["firing"] = firing
	SplitStatus["resolved"] = resolved
	return SplitStatus
	//return firing, resolved
	//golang template package requires functions to return just 1 argument, here is a related fragment (https://golang.org/src/text/template/funcs.go):
	//So, if you want 2 results, you probably need to create your own type and then return it
}

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		json := make(map[string]interface{})
		c.BindJSON(&json)
		funcMap := template.FuncMap{"ListSplitStatus": ListSplitStatus1}
		//t := template.Must(template.New("letter").Funcs(funcMap).Parse(letter))
		t, _ := template.New("template.tpl").Funcs(funcMap).ParseFiles("template.tpl")
		buf := bytes.NewBufferString("")
		t.Execute(buf, json)
		fmt.Println(buf.String())
		api_url := "https://oapi.dingtalk.com/robot/send?access_token=66a48ed58142ef88953fc67b794bc1b5d720f8e1a2a78fb6637c6485ce14dd5a"
		content := `{"msgtype": "markdown",
		"markdown": {
						"title": "prometheus",
						"text": "` + buf.String() + `"
		},
		"at": {
			"atMobiles": [
				"18301371817"
				],
				"isAtAll": false
		}
		}`
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
		c.JSON(200, gin.H{
			"receiver": json["receiver"],
			"status":   json["status"],
		})
	})
	router.Run(":8080")
	log.Println("message")
}
