package main

import (
	"go-mvc/app/utils"
	"go-mvc/app/utils/log"
	"go-mvc/routers"
	"os"

	_ "github.com/joho/godotenv/autoload" //自动加载.env文件
)

func main() {
	//初始化路由
	router := routers.BaseRouter{}.InitRouter()
	log.Instance().Info("启动服务")

	//初始化验证器
	utils.InitValidator()

	//启动http服务
	err := router.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}
