package main

import (
	"Test1/service"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type server struct {
	service.UnimplementedSayHelloServer
}

func (s server) SayHello(ctx context.Context, u *service.User) (*service.Addr, error) {
	fmt.Println(u.Age)
	return &service.Addr{City: "绵阳", House: u.Username}, nil
}

func main() {
	creds, _ := credentials.NewServerTLSFromFile("D:\\Go_Study\\GoPath\\src\\Test1\\protohello\\test.pem",
		"D:\\Go_Study\\GoPath\\src\\Test1\\protohello\\test.key")
	listen, _ := net.Listen("tcp", ":9090")
	//创建服务
	grpcServe := grpc.NewServer(grpc.Creds(creds))
	//注册服务
	service.RegisterSayHelloServer(grpcServe, &server{})
	//启动服务
	err := grpcServe.Serve(listen)
	if err != nil {
		fmt.Printf("fail:%v", err)
		return
	}
	user := &service.User{
		Username: "asas",
		Age:      18,
	}
	//序列化
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	//反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.String())
}
