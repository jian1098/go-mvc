package models

// Person结构体用于接收参数,如果是json数据把form改为json
// binding为参数验证器，更新规则可以参考下面两个网站
// 1.https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags
// 2.https://blog.csdn.net/weixin_40022980/article/details/122796567
type Person struct {
	Name string `form:"name"  binding:"required"`
	Age  int    `form:"age" binding:"required"`
}

// user表gorm结构体
type User struct {
	Id       int    `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;comment:ID" json:"id"` // 自增主键
	Name     string `gorm:"column:name;type:varchar(50);comment:用户名" json:"name"`
	Mobile   string `gorm:"column:mobile;type:varchar(100);comment:手机号码" json:"mobile"`
	Password string `gorm:"column:password;type:varchar(255);comment:密码" json:"password"`
}
