package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oceanlearn.teach/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//driverName := "mysql"
	//username := "root"
	username := viper.GetString("datasource.username")
	//password := "Root@123"
	password := viper.GetString("datasource.password")
	//host := "localhost"
	host := viper.GetString("datasource.host")
	//port := 3306
	port := viper.GetString("datasource.port")
	//database := "ginessential"
	database := viper.GetString("datasource.database")
	//charset := "utf8"
	charset := viper.GetString("datasource.charset")
	//timeout := "10s"
	timeout := viper.GetString("datasource.timeout")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s",
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
