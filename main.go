package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yaagOrSwaggerDemo/model"
	"yaagOrSwaggerDemo/siwag"
	siwag_gin "yaagOrSwaggerDemo/siwag/gin"
)

func main() {

	config := siwag.Config{
		IsOn: true,
		DocTitle: "Gin",
		DocPath: "apidoc.json",
		Author: "demo",
		Email: "1111@163.com",
		Host: "172.0.0.1",
		BasePath: "/",
		}

	//	init config file
	siwag.Init(&config)
	//add model
	siwag.AutoCreateJson(&model.User{},model.Permission{})

	r := gin.Default()
	// 调用中间件
	r.Use(siwag_gin.Document())
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "Hello World!"})
	})
	r.GET("/plain", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"result": "Hello World!"})
	})
	r.GET("/complex", func(c *gin.Context) {
		value := c.Query("key")
		c.JSON(http.StatusOK, gin.H{"value": value})
	})
	r.Run(":8081")
}
