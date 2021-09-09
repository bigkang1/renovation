package main

import (
	"captcha/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/mojocn/base64Captcha"
	"google.golang.org/grpc"
	"net"
)

//定义类
type CaptchaS struct {

}

var store = base64Captcha.DefaultMemStore

func (this *CaptchaS)GetCaptcha(ctx context.Context,c *pb.Captcha) (*pb.Captcha,error) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	a := base64Captcha.NewCaptcha(driver, store)

	var err error
	// 获取
	c.Vid, c.B64, err = a.Generate()
	return c,err
}

func (this *CaptchaS)CheckCaptcha(ctx context.Context,c *pb.CheckCap) (cr *pb.CheckCapRe,err error) {
	// 同时在内存清理掉这个图片
	re := store.Verify(c.Vid, c.Val, true)
	if !re {
		cr.Result = false
		return cr,err
	}
	cr.Result = true
	return cr,err
}

func main() {
	//初始化consul配置
	consulConfig := api.DefaultConfig()

	//创建consul对象
	consulClient ,err := api.NewClient(consulConfig)
	if err!=nil {
		fmt.Println("api.NewClient err:",err)
		return
	}

	//注册服务，服务的常规配置
	//告诉consul，即将注册的服务
	reg := api.AgentServiceRegistration{
		ID:"captcha",
		Tags:[]string{"captcha","test"},
		Name:"captcha",
		Address:"127.0.0.1",
		Port:1234,
		Check:&api.AgentServiceCheck{
			CheckID:"captcha test",
			TCP:"127.0.0.1:1234",
			Timeout:"1s",
			Interval:"5s",
		},
	}

	//注册grpc服务到consul
	consulClient.Agent().ServiceRegister(&reg)

	//初始化grpc对象
	grpcServer :=grpc.NewServer()

	//注册服务
	pb.RegisterCaptchaServer(grpcServer,new(CaptchaS))

	//设置监听，指定IP/port
	listener ,err :=net.Listen("tcp","127.0.0.1:1234")
	if err != nil{
		fmt.Println("Listen err",err)
	}
	defer listener.Close()

	fmt.Println("服务器启动.....")

	//启动服务
	grpcServer.Serve(listener)


}
