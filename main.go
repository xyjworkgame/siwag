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
		IsOn:     true,
		DocTitle: "Gin",
		DocPath:  "apidoc",
		Author:   "demo",
		Email:    "1111@163.com",
		Host:     "172.0.0.1",
		BasePath: "/",
	}
	//add model
	siwag.AutoCreateJson(model.User{}, model.Permission{})
	//	init config file
	siwag.Init(&config)


	r := gin.Default()
	// 调用中间件
	r.Use(siwag_gin.Document())
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "Hello World!"})
	})
	r.POST("/json", func(c *gin.Context) {
		form := c.PostForm("cesa")
		c.String(http.StatusOK, form)
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"result": "Hello World!"})
	})
	r.GET("/complex", func(c *gin.Context) {
		value := c.Query("key")
		c.JSON(http.StatusOK, gin.H{"value": value})
	})
	r.Run(":8080")
}
