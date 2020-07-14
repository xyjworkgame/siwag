package models

import (
	"github.com/go-openapi/spec"
)

// 基本数据,这些不用改变，
type Base struct {
	Swagger             string
	Info                *spec.Info
	Host                string
	BasePath            string
	Tags                *Tags
	Schemes             []string
	Paths               Paths
	SecurityDefinitions *spec.SecuritySchemeProps //这里不进行自动添加
	Definitions         map[string]*Definitions
}
type Definitions struct {
	Type       string
	Xml        map[string]string
	Properties map[string]interface{}
}

type Tags struct {
	Description string
	Name        string
}
