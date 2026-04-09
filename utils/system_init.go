package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	// 1. 设置配置文件的文件名（不带后缀）
	viper.SetConfigName("app")

	// 2. 设置配置文件的格式
	viper.SetConfigType("yml")

	// 3. 设置查找配置文件的路径
	viper.AddConfigPath("config")

	// 4. 正式开始读取文件
	err := viper.ReadInConfig()

	// 5. 错误处理
	if err != nil {
		fmt.Println("Error reading config file:", err)
	}

	fmt.Println("Config app:", viper.Get("app"))
	fmt.Println("Config mysql:", viper.Get("mysql"))
}
func InitMysql() {
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println("Mysql connected successfully, user:", user)
}
