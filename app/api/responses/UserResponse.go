package responses

/**
用户模块响应结构体
*/

// 用户信息
type UserInfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}
