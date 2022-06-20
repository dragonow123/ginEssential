package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oceanlearn.teach/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//driverName := "mysql"
	username := "root"
	password := "Root@123"
	host := "localhost"
	port := 3306
	database := "ginessential"
	charset := "utf8"
	timeout := "10s"
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		timeout)
	fmt.Println(args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err： " + err.Error())
	}
	defer db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
