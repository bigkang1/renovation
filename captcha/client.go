package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
	"captcha/pb"
)


func main()  {
	//初始化consul配置
	consulConfig := api.DefaultConfig()

	//创建consul对象(可以重新指定consul属性：IP/Port，也可以使用默认)
	consulClient,err := api.NewClient(consulConfig)

	//服务发现，从consul上，获取健康的服务
	services,_,err := consulClient.Health().Service("captcha","captcha",true,nil)
	if err!=nil {
		fmt.Println("err:",err)
		return
	}

	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	//连接服务
	//grpcConn,_ := grpc.Dial("127.0.0.1:8800",grpc.WithInsecure())

	//使用服务发现consul上的IP/Port来与服务建立连接
	grpcConn,_ := grpc.Dial(addr,grpc.WithInsecure())

	//初始化grpc客户端
	grpcClinent := pb.NewCaptchaClient(grpcConn)

	var captcha pb.Captcha

	//调用远程函数
	c,err := grpcClinent.GetCaptcha(context.TODO(),&captcha)

	fmt.Println(c,err)
}
