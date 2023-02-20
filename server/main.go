package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	service "grpc/proto" //这里之所以要这么写,就是因为包名和路径名不一致
	"net"
)

type server struct {
	service.UnimplementedSayHelloServer
}

// 方法重写
func (s *server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {

	//获取元数据信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}
	var appId string
	var appKey string
	//注意：如果在客户端中，appId和appKey是大写，这里必须小写才能行
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appId != "xiao" || appKey != "123" {
		return nil, errors.New("token不正确")
	}

	fmt.Println("server端 SayHello方法")
	return &service.HelloResponse{ResponseMsg: "hello" + req.GetRequestName()}, nil
}

func main() {
	//TSL认证，有两个参数cretFile，keyFile
	//我们需要将刚才生成的自签名证书test.pem 和 私钥文件 test.key 放进去即可
	//creds, _ := credentials.NewServerTLSFromFile("D:\\学习笔记\\后端学习资料\\grpc\\key\\test.pem",
	//	"D:\\学习笔记\\后端学习资料\\grpc\\key\\test.key")

	//开启端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	//创建grpc服务,将服务暴露供别人使用
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	//在grpc服务端中注册我们自己编写的服务
	//这里必须通过引用注册,即必须是&server{},因为上面方法实现是通过指针实现的,而RegisterSayHelloServer第二个参数
	//是接口类型,指针实现的方法只能传入指针类型才能得到该方法
	service.RegisterSayHelloServer(grpcServer, &server{}) //注册服务和方法

	//启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("grpcServer failed,err:", err)
		return
	}
}
