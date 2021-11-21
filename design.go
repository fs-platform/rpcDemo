package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Base interface {
	Say(messages string, reply *string) error
}

func RegisterService(Base Base) {
	err := rpc.RegisterName("People", Base)
	if err != nil {
		fmt.Printf("注册服务失败%s", err.Error())
	}
}

func InitServer() net.Conn {
	//监听端口号
	listener, err := net.Listen("tcp", "127.0.0.1:8083")
	if err != nil {
		fmt.Printf("监听端口失败%s", err.Error())
	}
	//等待客户端请求
	fmt.Println("等待客户端请求")
	con, err := listener.Accept()
	if err != nil {
		fmt.Printf("接受客户端请求异常%s", err.Error())
	}
	//绑定rpc服务
	jsonrpc.ServeConn(con)
	defer func() {
		con.Close()
		listener.Close()
	}()
	return con
}

type client struct {
	Con *rpc.Client
}

func initClient(addr string) client {
	con, _ := jsonrpc.Dial("tcp", addr)
	return client{Con: con}
}

func (c client) Say(message string, reply *string) error {
	return c.Con.Call("Peoples.Say", message, reply)
}
