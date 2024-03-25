package main

import (
	"github.com/joho/godotenv"
	"go-mvc/routers"
	"log"
	"os"
)

func main() {
	//获取环境变量
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//启动http服务
	router := routers.BaseRouter{}.InitRouter()
	err = router.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}
