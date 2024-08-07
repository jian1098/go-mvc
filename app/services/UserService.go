package services

import (
	"go-mvc/app/api/responses"
	"go-mvc/app/models"
	"go-mvc/app/utils/db"
)

// 用户服务
type UserService struct {
	BaseService //继承BaseService
}

func (service UserService) GetUserInfo(userId int) (*responses.UserInfo, error) {
	//查找用户
	var user models.User
	var userInfo responses.UserInfo
	err := db.Instance().Where("id = ?", userId).Select("id, name, mobile").First(&user).Error
	if err != nil {
		return &userInfo, err
	}

	//用户信息
	userInfo.Id = user.Id
	userInfo.Name = user.Name
	userInfo.Mobile = user.Mobile

	return &userInfo, nil
}
