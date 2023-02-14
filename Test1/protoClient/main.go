package main

import (
	"Test1/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//不加密
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v\n", err)
	}
	defer conn.Close()
	client := service.NewSayHelloClient(conn)
	ad, _ := client.SayHello(context.Background(), &service.User{Username: "renwei", Age: 24})
	fmt.Println(*ad)
}
