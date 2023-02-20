package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	service "grpc/proto"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "xiao",
		"appkey": "123",
	}, nil
}
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {

	//creds, _ := credentials.NewClientTLSFromFile("D:\\学习笔记\\后端学习资料\\grpc\\key\\test.pem",
	//	"*.xiaokk.com")

	//连接到server端,此处禁用了安全传输,没有加密和验证,insecure.NewCredentials()不开启安全传输

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		fmt.Printf("connect failed,err:%v\n", err)
		return
	}
	defer conn.Close()

	//建立连接
	client := service.NewSayHelloClient(conn)

	//执行rpc调用(这个方法在服务器端来实现并返回结果),调用的SayHello方法是服务端实现的方法(即刚才服务端重写的方法)
	resp, err := client.SayHello(context.Background(), &service.HelloRequest{RequestName: "张瑜"})
	if err != nil {
		fmt.Println("client.SayHello failed,err:", err)
		fmt.Println("hello git!")
		fmt.Println("hello git hot-fix")
		fmt.Println("master test 1")
		fmt.Println("hot-fix test1")
		fmt.Println("push test")
		fmt.Println("push test1")
		fmt.Println("pull test")
		return
	}
	//获取回复的消息
	fmt.Println(resp.GetResponseMsg())
}
