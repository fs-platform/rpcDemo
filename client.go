package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	var reply = ""
	con, err := rpc.Dial("tcp", "127.0.0.1:8083")
	message := "hello"
	if err != nil {
		fmt.Printf("远程连接失败%s", err.Error())
	}
	err = con.Call("People.Say", message, &reply)
	if err != nil {
		fmt.Printf("数据返回异常%s", err.Error())
	}
	defer con.Close()
	fmt.Println(reply)
}
