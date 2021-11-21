package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"rpc/pb"
)

type ShowTimeClient struct {
}

func main() {
	var opts []grpc.DialOption
	serverAddr := flag.String("addr", "localhost:8888", "The server address in the format of host:port")
	opts = append(opts, grpc.WithInsecure())
	cc, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		fmt.Printf("错误:%s", err)
		return
	}
	client := pb.NewServerTimeClient(cc)
	inParam := new(pb.TimeRequire)
	inParam.Type = "second"
	TimeInfo, _ := client.ShowTime(context.Background(), inParam)
	fmt.Println(TimeInfo)
}
