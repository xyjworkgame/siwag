package gin

import (
	"github.com/betacraft/yaag/yaag"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"yaagOrSwaggerDemo/middleware"
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

		swaggerCall := model.SwaggerSpec{}
		middleware.Before(&swaggerCall, c.Request)
		c.Next()
		if yaag.IsStatusCodeValid(c.Writer.Status()) {

			headers := map[string]string{}
			for k, v := range c.Writer.Header() {
				log.Println(k, v)
				headers[k] = strings.Join(v, " ")
			}

			//go yaag.GenerateHtml(&apiCall)
		}
	}
}