package models

import (
	"github.com/go-openapi/spec"
)

//type SwaggerProps struct {
//	ID                  string                 `json:"id,omitempty"`
//	Consumes            []string               `json:"consumes,omitempty"`
//	Produces            []string               `json:"produces,omitempty"`
//	Schemes             []string               `json:"schemes,omitempty"`
//	Swagger             string                 `json:"swagger,omitempty"`
//	Info                *Info                  `json:"info,omitempty"`
//	Host                string                 `json:"host,omitempty"`
//	BasePath            string                 `json:"basePath,omitempty"`
//	Paths               *Paths                 `json:"paths"`
//	Definitions         Definitions            `json:"definitions,omitempty"`
//	Parameters          map[string]Parameter   `json:"parameters,omitempty"`
//	Responses           map[string]Response    `json:"responses,omitempty"`
//	SecurityDefinitions SecurityDefinitions    `json:"securityDefinitions,omitempty"`
//	Security            []map[string][]string  `json:"security,omitempty"`
//	Tags                []Tag                  `json:"tags,omitempty"`
//	ExternalDocs        *ExternalDocumentation `json:"externalDocs,omitempty"`
//}
// 基本数据
type Base struct {
	Swagger string
	Info  *spec.Info
	Host string
	BasePath string
	Tags *spec.TagProps
	Schemes []string
	Paths map[string]Paths
	SecurityDefinitions *spec.SecuritySchemeProps
	Definitions map[string]*Definitions
}
type Definitions struct {
	Type string
	Xml map[string]string
	Properties map[string]interface{}
}

// path 下的详细信息
type PathItemProps struct {
	Get        *spec.OperationProps
	Put        *spec.OperationProps
	Post       *spec.OperationProps
	Delete     *spec.OperationProps
	Options    *spec.OperationProps
	Head       *spec.OperationProps
	Patch      *spec.OperationProps
	Parameters []spec.ParamProps
}


