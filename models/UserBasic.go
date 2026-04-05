package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string // 用户名
	Password      string // 密码
	Phone         string // 电话
	Email         string // 邮箱
	Identity      string // 身份
	ClientIp      string // 客户端IP
	ClientPort    string // 客户端端口
	LoginTime     string // 登录时间
	HeartbeatTime string // 心跳时间
	LogOutTime    string // 登出时间
	IsLogOut      bool   // 是否登出
	DeviceInfo    string // 设备信息
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
