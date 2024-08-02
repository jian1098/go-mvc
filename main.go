package main

import (
	"go-mvc/app/utils"
	"go-mvc/app/utils/db"
	"go-mvc/routers"
	"os"

	_ "github.com/joho/godotenv/autoload" //自动加载.env文件
)

func main() {
	//初始化路由
	router := routers.BaseRouter{}.InitRouter()

	//初始化验证器
	utils.InitValidator()

	//初始化数据库连接
	db.InitConn()

	//启动http服务
	err := router.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}
