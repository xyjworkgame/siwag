package middleware

import (
	"github.com/betacraft/yaag/middleware"
	"github.com/betacraft/yaag/yaag"
	"github.com/betacraft/yaag/yaag/models"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)
// 每次记录 请求的的前后数据，记录
func Document() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !yaag.IsOn() {
			return
		}
		apiCall := models.ApiCall{}
		middleware.Before(&apiCall, c.Request)
		//blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		//c.Writer = blw
		c.Next()


		if yaag.IsStatusCodeValid(c.Writer.Status()) {
			apiCall.MethodType = c.Request.Method
			apiCall.CurrentPath = strings.Split(c.Request.RequestURI, "?")[0]

			apiCall.ResponseBody = GinBodyLogMiddleware(c)
			apiCall.ResponseCode = c.Writer.Status()
			headers := map[string]string{}
			for k, v := range c.Writer.Header() {
				log.Println(k, v)
				headers[k] = strings.Join(v, " ")
			}
			apiCall.ResponseHeader = headers
			//go format.GenerateJson()
		}
	}
}
