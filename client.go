package main

import "fmt"

func main() {
	var reply = ""
	//rpc 建立连接
	client := initClient("127.0.0.1:8083")

	err := client.Say("你好", &reply)
	if err != nil {
		fmt.Printf("say err %s", err.Error())
	}
	fmt.Println(reply)
}
