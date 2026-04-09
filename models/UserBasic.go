package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    // 用户名
	Password      string    // 密码
	Phone         string    // 电话
	Email         string    // 邮箱
	Identity      string    // 身份
	ClientIp      string    // 客户端IP
	ClientPort    string    // 客户端端口
	LoginTime     time.Time // 登录时间
	HeartbeatTime time.Time // 心跳时间
	LoginOutTime  time.Time // 登出时间
	IsLogOut      bool      // 是否登出
	DeviceInfo    string    // 设备信息
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserBasicList() []*UserBasic {
	data := make([]*UserBasic, 0)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println("UserBasic:", v)
	}
	return data
}
