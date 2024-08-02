<div align="center">
<br/>
<br/>
  <h1 align="center">
    Go-MVC
  </h1>
</div>

项目地址：https://github.com/jian1098/go-mvc



## 项目说明

go语言gin框架搭建的MVC模式（参考PHP语言Laravel和ThinkPHP框架结构）web开发框架项目，集成了数据库gorm、类型转换cast等常用包，以及一些代码示例；控制器层和路由层做了仿继承处理，方便添加统一的属性和方法，模型、模板、控制器、路由、常量、验证器、服务、全局日志、全局DB等都做了分离处理，分配到各个模块



## 运行环境

本项目测试环境为`go 1.22.0`，理论上` go version >= 1.20`即可



## 目录结构

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
|   |   `-- requests						 #api请求参数结构体
|   |   `-- responses						 #api响应结构体
|   |-- utils								#公共模块
|   |   |-- db.go
|   |   `-- zapLogger.go
|   |-- constants							 #常量目录
|   |   `-- Response.go
|   |-- middlewares							 #中间件目录
|   |   `-- Cors.go
|   |-- services							 #服务层目录
|   |   `-- UserServices.go
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





## 如何使用

### 下载源码

```bash
git clone https://github.com/jian1098/go-mvc.git
```

创建配置文件，并修改为你的配置

```bash
cd go-mvc
cp .env.example .env
vim .env
```



### 开启go mod

```bash
go env -w GO111MODULE=on
```



### 设置代理

```bash
go env -w GOPROXY="https://goproxy.cn"
```



### 安装依赖

```bash
go mod vendor
```



### 运行http服务

```bash
go run main.go
```



### 添加路由

在`routers`目录下对应模块增加路由



### 访问API接口

请求路由以`/api/`开头， 例如：`127.0.0.1:8080/api/index/demo`



### 访问后台页面

请求路由以`/admin/`开头， 例如：`127.0.0.1:8080/admin/home/index`



## 第三方包

| 包名                           | 用途说明                   |
| ------------------------------ | -------------------------- |
| github.com/gin-gonic/gin       | gin框架                    |
| github.com/joho/godotenv       | 加载.env配置               |
| go.uber.org/zap                | zap日志                    |
| github.com/spf13/cast          | 变量类型转换               |
| github.com/jinzhu/gorm         | gorm数据库包               |
| github.com/go-sql-driver/mysql | mysql驱动，配合gorm使用    |
| github.com/gookit/validate     | 验证器，支持验证场景和标签 |
| github.com/dgrijalva/jwt-go    | jwt                        |

