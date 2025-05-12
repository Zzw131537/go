/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-06 14:50:34
 */
package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"` // 姓名
	Email          string // 邮箱
	PasswordDigest string // 加密后的密码
	NickName       string // 昵称
	Status         string // 状态
	Avatar         string // 头像
	Money          string // 现金
}

const (
	PassWordCost        = 12 // 密码加密难度
	Active       string = "active"
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
