package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"yaagOrSwaggerDemo/middleware"
	model "yaagOrSwaggerDemo/siwag/models"
)

/*
@Time : 2020/7/14 23:32
@Author : Firewine
@File : siwagmiddleware
@Software: GoLand
@Description:
*/
const MaxInMemoryMultipartSize = 32000000

func After(siwagCall *model.Path, c *gin.Context) {

	headers := middleware.ReadHeaders(c.Request)
	// 3. 判断body 里面的参数数据  ,应该在after里面，next前面无法获取到
	val, ok := headers["Content-Type"]
	log.Println(val)
	if ok {
		ct := strings.TrimSpace(headers["Content-Type"])
		switch ct {
		case "application/x-www-form-urlencoded":
			fallthrough
		case "application/json, application/x-www-form-urlencoded":
			log.Println("Reading form")
			bodyResult := ReadBodyParams(c)
			for k, v := range bodyResult {
				parameter := model.Parameter{}
				parameter.Name = k
				parameter.In = "path"
				parameter.Required = true
				parameter.Type = reflect.TypeOf(v).String()
				jsonStr, _ := json.Marshal(bodyResult)
				parameter.Schema.Example = string(jsonStr)
				//log.Println(parameter.Type)

				siwagCall.Parameters = append(siwagCall.Parameters, parameter)
			}
		case "application/json":
			log.Println("Reading body")

			body := *middleware.ReadBody(c.Request)
			var mapResult map[string]interface{}
			if err := json.Unmarshal([]byte(body), &mapResult); err != nil {
				log.Println(err)
				return
			}
			for k, v := range mapResult {
				parameter := model.Parameter{}
				parameter.In = "body"
				parameter.Required = true
				parameter.Name = k
				parameter.Type = reflect.TypeOf(v).String()

				siwagCall.Parameters = append(siwagCall.Parameters, parameter)
			}

		default:
			if strings.Contains(ct, "multipart/form-data") {

				handleMultipart(siwagCall, c.Request)

			} else {
				log.Println("Reading body")
				log.Println(*middleware.ReadBody(c.Request))

			}
		}
	}

	siwagCall.Consumes = []string{
		c.Request.Header.Get("Content-Type"),
	}
}

func Before(siwagCall *model.Path, c *gin.Context) {

	//	这里就是判断以下数据
	//	1. 参数 query，path ，body{form，json，and so on}
	queryResult := ReadQueryParams(c.Request)
	// 遍历添加数据
	for k, v := range queryResult {
		parameter := model.Parameter{}
		parameter.Name = k
		parameter.In = "query"
		parameter.Required = true
		//	TODO 这里有可能无法实现添加参数类型，因为得到的都是字符串，除非通过转换切换
		//log.Println(reflect.ValueOf(v).String())
		parameter.Type = reflect.TypeOf(v).String()
		jsonStr, _ := json.Marshal(queryResult)
		parameter.Schema.Example = string(jsonStr)
		// 添加进去
		siwagCall.Parameters = append(siwagCall.Parameters, parameter)
	}

	// 2. 判断path 路径上面的参数 ,path 参数在context 里面，这里没法写
	pathResult := ReadPathParams(c)
	for k, v := range pathResult {
		parameter := model.Parameter{}
		parameter.Name = k
		parameter.In = "path"
		parameter.Required = true
		parameter.Type = reflect.ValueOf(v).String()
		//log.Println(parameter.Type)

		siwagCall.Parameters = append(siwagCall.Parameters, parameter)
	}

}

func ReadBodyParams(c *gin.Context) map[string]string {
	forms := c.Request.Form
	results := map[string]string{}
	for k, v := range forms {
		results[k] = v[0]

	}
	return results
}

func ReadPathParams(c *gin.Context) map[string]string {

	results := map[string]string{}
	params := c.Params // path上面的参数
	for _, v := range params {
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


func handleMultipart(siwagCall *model.Path, req *http.Request) {

	if err := req.ParseMultipartForm(MaxInMemoryMultipartSize); err != nil {
		log.Println(err)
		return
	}
	postForm := ReadMultiPostForm(req.MultipartForm)
	for k, _ := range postForm {
		parameter := model.Parameter{}

		parameter.Name = k
		parameter.Required = true
		parameter.In = "body"

		//jsonStr, _ := json.Marshal(map[string]string{
		//	k:v,
		//})
		//parameter.Schema.Example = ""

		siwagCall.Parameters = append(siwagCall.Parameters, parameter)
	}

}

func ReadMultiPostForm(mpForm *multipart.Form) map[string]string {
	postForm := map[string]string{}
	for key, val := range mpForm.File {
		postForm[key] = val[0].Header.Get("Content-Type")
	}
	for key, val := range mpForm.Value {
		postForm[key] = val[0]
	}
	return postForm
}
