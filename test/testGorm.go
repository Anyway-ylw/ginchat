package main

import (
	"fmt"
	"ginchat/models"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Just test gorm")
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})
	// create
	user := &models.UserBasic{
		Name:          "test",
		LoginTime:     time.Now(), // 必须赋值
		HeartbeatTime: time.Now(),
		LoginOutTime:  time.Now(),
	}
	db.Create(&user)

	// Read
	fmt.Println(db.First(user, 1))

	db.Model(user).Update("PassWord", "123456")

}
