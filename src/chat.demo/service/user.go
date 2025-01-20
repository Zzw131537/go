package service

import (
	"chat/model"
	"chat/serializer"
)

type UserRegisterService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
}

func (service *UserRegisterService) Register() serializer.Response {
	var user model.User
	count := 0
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count != 0 { // 用户名已经注册
		return serializer.Response{
			Status: 400,
			Msg:    "用户名已经存在了",
		}
	}
	user = model.User{
		UserName: service.UserName,
	}
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "加密出错",
		}
	}
	model.DB.Create(&user)
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}
