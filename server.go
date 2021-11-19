package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type People struct {
}

// Say 1.方法必须包外可见 2.方法必须有个参数,都是导出类型,内建类型 3.第二个参数必须是指针 4. 方法只能返回一个error
func (This People) Say(messages string, reply *string) error {
	*reply = "Aron say:" + messages
	return nil
}

func main() {
	//注册服务名称
	err := rpc.RegisterName("People", new(People))
	if err != nil {
		fmt.Printf("注册服务失败%s", err.Error())
	}
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
	rpc.ServeConn(con)

	defer func() {
		con.Close()
		listener.Close()
	}()
}
