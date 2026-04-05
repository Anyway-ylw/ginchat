package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("test gorm")
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})
	// create
	user := &models.UserBasic{
		Name: "test",
	}
	db.Create(&user)

	// Read
	fmt.Println(db.First(user, 1))

	db.Model(user).Update("PassWord", "123456")

}
