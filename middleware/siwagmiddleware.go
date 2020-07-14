package middleware

import (
	"github.com/go-openapi/spec"
	"net/http"
	"net/url"
	model "yaagOrSwaggerDemo/siwag/models"
)

/*
@Time : 2020/7/14 23:32
@Author : Firewine
@File : siwagmiddleware
@Software: GoLand
@Description:
*/

func Before(apiCall *model.Path, req *http.Request) {
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
	var parameter model.Parameter
//	这里就是判断以下数据
//	1. 参数 query，path ，body{form，json，and so on}
	queryResult := ReadQueryParams(req)
	// 遍历添加数据
	for k,v := range queryResult{
		parameter.Name = k
		parameter.In = "query"
		parameter.Required = true
	//	TODO 这里实现添加参数例子，
	}


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