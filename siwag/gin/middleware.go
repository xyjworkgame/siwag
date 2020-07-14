package gin

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"yaagOrSwaggerDemo/middleware"
	"yaagOrSwaggerDemo/siwag"
	model "yaagOrSwaggerDemo/siwag/models"
)

/*
@Time : 2020/7/12 15:40
@Author : Firewine
@File : server
@Software: GoLand
@Description:
*/
// 数据获取
func Document() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !siwag.IsOn() {
			return
		}
		var siwagCalls model.Paths

		var siwagCallItems model.PathItems

		siwagCall := model.Path{}
		middleware.Before(&siwagCall, c.Request)

		c.Next()

		// 获取响应的数据
		if siwag.IsStatusCodeValid(c.Writer.Status()) {

			headers := map[string]string{}
			for k, v := range c.Writer.Header() {
				log.Println(k, v)
				headers[k] = strings.Join(v, " ")
			}

			// 填补数据(没请求一个，就写入文件里面)
			// 获取请求后的一些信息
			// 设置url 为key
			siwagCallItems[c.Request.Method] = siwagCall
			siwagCalls[c.Request.URL.String()] = siwagCallItems
			// 存储文件里面
			siwag.InitInfo.Paths = siwagCalls
			go siwag.GenerateJson(&siwag.InitInfo)
		}

	}
}
