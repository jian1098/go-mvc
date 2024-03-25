package structs

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
	Id       uint   `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT;comment:ID" json:"id"` // 自增主键
	Nickname string `gorm:"column:nickname;type:varchar(50);comment:昵称" json:"nickname"`
	Email    string `gorm:"column:email;type:varchar(100);comment:邮箱" json:"email"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`
}
