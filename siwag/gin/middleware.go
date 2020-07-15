package gin

import (
	"github.com/gin-gonic/gin"
	"yaagOrSwaggerDemo/siwag"
	middleware "yaagOrSwaggerDemo/siwag/gin/middleware"
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
		siwagPaths := model.Paths{}
		siwagPathItems:= model.PathItems{}
		siwagPath := model.Path{}
		siwagResponses := model.Responses{}
		//siwagResponse := model.Response{}
		middleware.Before(&siwagPath, c)
		c.Next()
		middleware.After(&siwagPath,c)
		// 获取响应的数据
		if siwag.IsStatusCodeValid(c.Writer.Status()) {

			//headers := map[string]string{}
			//for k, v := range c.Writer.Header() {
			//	log.Println(k, v)
			//	headers[k] = strings.Join(v, " ")
			//}
			siwagPath.Produces = []string{
				"application/json",
			}
			siwagResponses.StatusCodeResponse = map[int]model.Response{
				c.Writer.Status(): model.Response{},
			}

			siwagPath.Response = &siwagResponses
			siwagPathItems[c.Request.Method] = siwagPath
			siwagPaths[c.FullPath()] = siwagPathItems
			// 存储文件里面
			siwag.InitInfo.Paths = siwagPaths
			go siwag.GenerateJson(&siwag.InitInfo)
		}

	}
}
