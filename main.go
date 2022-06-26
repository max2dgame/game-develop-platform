package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
)

func main() {
	startServer()
}

func startServer() {
	//gin.SetMode(gin.DebugMode)
	//gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	//init engine

	//init config

	r.Static("/uploads", "./uploads")
	r.Static("/max2d", "./client")
	r.GET("/", func(c *gin.Context) {
		//跳转到/luyou2对应的路由处理函数
		c.Request.URL.Path = "/max2d" //把请求的URL修改
		r.HandleContext(c)            //继续后续处理
	})

	//init model

	_ = r.Run(":8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	//eng.MysqlConnection().Close()
}
