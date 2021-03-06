package gin

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"strings"
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
		siwagPathItems := model.PathItems{}
		siwagPath := model.Path{}
		//siwagResponse := model.Response{}
		middleware.Before(&siwagPath, c)
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		middleware.After(&siwagPath, c)
		// 获取响应的数据
		if siwag.IsStatusCodeValid(c.Writer.Status()) {
			siwagPath.Produces = []string{
				"application/json",
			}
			var tags []model.Tags
			tags = append(tags,model.Tags{
				Name:"default",
				})
			siwagPath.Tags = []string{"default"}
			response := model.Response{
				Description: blw.body.String(),
			}
			responsec := map[int]model.Response{
				c.Writer.Status(): response,
			}
			siwagPath.Responses = responsec

			siwagPathItems[strings.ToLower(c.Request.Method)] = siwagPath

			if c.FullPath() == "" {
				siwagPaths[c.Request.URL.Path] = siwagPathItems

			} else {
				siwagPaths[c.FullPath()] = siwagPathItems

			}
			siwagPath.Schemes = []string{"http","https"}
			siwag.InitInfo.Tags = &tags

			// 存储文件里面
			siwag.InitInfo.Paths = siwagPaths
			go siwag.GenerateJson(&siwag.InitInfo)
		}

	}
}
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinBodyLogMiddleware(c *gin.Context) string {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	//statusCode := c.Writer.Status()
	//if statusCode >= 200 {
	//	//ok this is an request with error, let's make a record for it
	//	// now print body (or log in your preferred way)
	//	fmt.Println("Response body: " + blw.body.String())
	//}
	return blw.body.String()
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
