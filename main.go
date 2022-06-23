package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"oceanlearn.teach/ginessential/common"
	"os"
)

func main() {
	InitConfig()
	common.InitDB()
	//defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
