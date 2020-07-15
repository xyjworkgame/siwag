package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"reflect"
	model "yaagOrSwaggerDemo/siwag/models"
)

/*
@Time : 2020/7/14 23:32
@Author : Firewine
@File : siwagmiddleware
@Software: GoLand
@Description:
*/
func After(siwagCall *model.Path,c *gin.Context){

}
func Before(siwagCall *model.Path, c *gin.Context) {
	/*
		Description string   // 描述
		Consumes    []string //
		Produces    []string //
		Schemes     []string
		Tags        []string //手动添加
		Summary     string
		ID          string
		Deprecated  bool
		Parameters  []Parameter
		Response    *Responses
	*/

	//	这里就是判断以下数据
	//	1. 参数 query，path ，body{form，json，and so on}
	queryResult := ReadQueryParams(c.Request)
	// 遍历添加数据
	for k, v := range queryResult {
		var parameter model.Parameter
		parameter.Name = k
		parameter.In = "query"
		parameter.Required = true
		//	TODO 这里有可能无法实现添加参数类型，因为得到的都是字符串，除非通过转换切换
		log.Println(reflect.ValueOf(v).String())
		parameter.Type = reflect.TypeOf(v).String()
		// 添加进去
		siwagCall.Parameters = append(siwagCall.Parameters, parameter)
	}

	// 2. 判断path 路径上面的参数 ,path 参数在context 里面，这里没法写
	pathResult := ReadPathParams(c)
	for k,v := range pathResult{
		var parameter model.Parameter
		parameter.Name = k
		parameter.In = "path"
		parameter.Required= true
		parameter.Type = reflect.ValueOf(v).String()
		log.Println(parameter.Type)

		siwagCall.Parameters = append(siwagCall.Parameters,parameter)
	}
	// 3. 判断body 里面的参数数据  ,应该在after里面，next前面无法获取到
	//bodyResult := ReadBodyParams(c)
	//fmt.Println(bodyResult)

}

func ReadBodyParams(c *gin.Context) map[string]string {
	forms := c.Request.Form
	results := map[string]string{}
	for k,v := range forms{
		log.Println(k,v)
	}
	return results
}

func ReadPathParams(c *gin.Context) map[string]string {

	results := map[string]string{}
	params := c.Params // path上面的参数
	for _,v := range params{
		results[v.Key] = v.Value

	}
	return results
}

func ReadQueryParams(req *http.Request) map[string]string {
	params := map[string]string{}
	u, err := url.Parse(req.RequestURI)
	if err != nil {
		return params
	}
	for k, v := range u.Query() {
		if len(v) < 1 {
			continue
		}
		params[k] = v[0]
	}
	return params
}
