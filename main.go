package main

import (
	"github.com/gin-gonic/gin"
	"oceanlearn.teach/ginessential/common"
)

func main() {
	common.InitDB()
	//defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
