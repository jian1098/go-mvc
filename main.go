package main

import (
	_ "github.com/joho/godotenv/autoload" //自动加载.env文件
	"go-mvc/routers"
	"os"
)

func main() {
	//启动http服务
	router := routers.BaseRouter{}.InitRouter()
	err := router.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}
