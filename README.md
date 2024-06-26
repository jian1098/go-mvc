# go-mvc
go语言gin框架搭建的MVC模式web项目，集成了数据库gorm、类型转换cast等常用包，以及一些代码示例；控制器层和路由层做了仿继承处理，方便添加统一的属性和方法，模型、模板、控制器、路由、常量、模块等都做了分离处理，分配到各个模块

本项目测试环境为`go 1.22.0`



### 目录结构

```
|-- app										#开发主目录
|   |-- admin								#后台模块
|   |   `-- controllers						 #后台控制器
|   |       |-- adminController.go
|   |       |-- baseController.go
|   |       `-- homeController.go
|   |-- api									#api模块
|   |   `-- controllers						 #api控制器
|   |       |-- baseController.go
|   |       `-- indexController.go
|   |-- common								#公共模块
|   |   |-- db.go
|   |   `-- zapLogger.go
|   |-- constants							 #常量目录
|   |   `-- Response.go
|   |-- middlewares							 #中间件目录
|   |   `-- IpBlackList.go
|   `-- models								 #模型结构体目录
|       `-- User.go
|-- .env									#环境配置
|-- go.mod
|-- go.sum
|-- main.go									#main.go
|-- routers									#路由模块
|   |-- adminRouter.go
|   |-- apiRouter.go
|   `-- baseRouter.go
|-- static									#静态文件
|   `-- js
|       `-- jquery.min.js
|-- templates								#模板文件
|    `-- admin
|       |-- home
|       |   `-- index.html
|       `-- index
|           `-- index.html
|-- runtimes								
|	 `-- log								#日志文件目录
|-- uploads								   	#上传文件目录
```





### 如何使用

下载源码

```bash
git clone https://github.com/jian1098/go-mvc.git
```

创建配置文件，并修改为你的配置

```bash
cd go-mvc
cp .env.example .env
vim .env
```

开启go mod

```bash
go env -w GO111MODULE=on
```

设置代理

```bash
go env -w GOPROXY="https://goproxy.cn"
```

安装依赖

```bash
go mod vendor
```

运行http服务

```bash
go run main.go
```





### 第三方包

| 包名                           | 用途说明                |
| ------------------------------ | ----------------------- |
| github.com/gin-gonic/gin       | gin框架                 |
| github.com/joho/godotenv       | 加载.env配置            |
| go.uber.org/zap                | zap日志                 |
| github.com/spf13/cast          | 变量类型转换            |
| github.com/jinzhu/gorm         | gorm数据库包            |
| github.com/go-sql-driver/mysql | mysql驱动，配合gorm使用 |
|                                |                         |

