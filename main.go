package main

import (
	_ "github.com/joho/godotenv/autoload" //自动加载.env文件
	"go-mvc/app/common"
	"go-mvc/routers"
	"os"
)

func main() {
	//初始化路由
	router := routers.BaseRouter{}.InitRouter()

	//初始化验证器
	common.InitValidator()

	//启动http服务
	err := router.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}
