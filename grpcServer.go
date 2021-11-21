package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rpc/pb"
	"time"
)

type serverTime struct {
}

func (serverTime) ShowTime(ctx context.Context, timeParam *pb.TimeRequire) (*pb.TimeInfo, error) {
	result := new(pb.TimeInfo)
	if timeParam.Type == "UTC" {
		result.TimeUTC = time.Now().Unix()
	} else if timeParam.Type == "second" {
		result.TimeSecond = time.Now().UnixNano()
	} else {
		err := errors.New("无效type")
		return nil, err
	}

	return result, nil
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8888")
	grpcServer := grpc.NewServer()
	pb.RegisterServerTimeServer(grpcServer, &serverTime{})
	fmt.Println("正在等待请求")
	grpcServer.Serve(listener)
}
