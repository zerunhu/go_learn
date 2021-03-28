package main

import "fmt"

func main() {
	content := `{"msgtype": "markdown",
		"markdown": {
						"title": "prometheus",
						"text": "d"
		},
		"at": {
			"atMobiles": [
				"18301371817"
				],
				"isAtAll": false
		}
		}`
	c := "def"
	fmt.Println(content)
	fmt.Printf("abc%v", c)
}
