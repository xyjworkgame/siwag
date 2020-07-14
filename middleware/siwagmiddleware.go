package middleware

import (
	"github.com/go-openapi/spec"
	"net/http"
	"net/url"
	"reflect"
)

/* 32 MB in memory max */
//const MaxInMemoryMultipartSize = 32000000

var reqWriteExcludeHeaderDump = map[string]bool{
	"Host":              true, // not in Header map anyway
	"Accept":            false,
	"Content-Type":      true,
	"Content-Length":    false,
	"Transfer-Encoding": false,
	"Trailer":           false,
	"Accept-Encoding":   false,
	"Accept-Language":   false,
	"Cache-Control":     false,
	"Connection":        false,
	"Origin":            false,
	"User-Agent":        false,
}

// 查找请求头，请求参数
func Before(siwagCall *spec.OperationProps, req *http.Request) {
	var parameters []spec.Parameter

	//query
	//body
	//paths
	//
	parameters = ReadQueryParams(req)
	siwagCall.R
}

// 获取 query 类型 的参数
func ReadQueryParams(req *http.Request) []spec.Parameter {
	var parameter []spec.Parameter

	u, err := url.Parse(req.RequestURI)
	if err != nil {
		return parameter
	}
	for k, v := range u.Query() {
		params := spec.Parameter{}
		if len(v) < 1 {
			continue
		}

		params.Name = k
		params.In = "query"
		params.Description = ""
		params.Required = true //默认都是true
		params.Type = reflect.ValueOf(v).String()
		parameter = append(parameter, params)

	}
	return parameter
}
