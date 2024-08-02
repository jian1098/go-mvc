package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var conn *gorm.DB

/*
*
DB实例
*/
func Instance() *gorm.DB {
	if conn == nil {
		conn = InitConn()
	}
	return conn
}

/**
 * 初始化数据库连接
 */
func InitConn() *gorm.DB {
	//获取环境变量
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_HOSTPORT")
	database := os.Getenv("DB_DATABASE")
	charset := os.Getenv("DB_CHARSET")
	fmt.Println(username + ":" + password + "@(" + host + ":" + port + ")/" + database + "?charset=" + charset + "&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", username+":"+password+"@("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return os.Getenv("DB_PREFIX") + defaultTableName
	}

	//是否禁用表名自动加s后缀
	db.SingularTable(cast.ToBool(os.Getenv("DB_SINGULAR_TABLE")))

	//连接池设置
	db.DB().SetMaxIdleConns(cast.ToInt(os.Getenv("DB_MAX_IDLE_CONNS"))) //最大空闲连接数
	db.DB().SetMaxOpenConns(cast.ToInt(os.Getenv("DB_MAX_OPEN_CONNS"))) //最大连接数

	return db
}
