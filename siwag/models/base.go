package models

import (
	"github.com/go-openapi/spec"
)

// 基本数据,这些不用改变，
type Base struct {
	Swagger             string                    `json:"swagger"`
	Info                *spec.Info                `json:"info"`
	Host                string                    `json:"host"`
	BasePath            string                    `json:"basePath"`
	Tags                *[]Tags                   `json:"tags"`
	Schemes             []string                  `json:"schemes"`
	Paths               Paths                     `json:"paths"`
	SecurityDefinitions *spec.SecuritySchemeProps `json:"securityDefinitions"`
	Definitions         map[string]*Definitions   `json:"definitions"`
}
type Definitions struct {
	Type string `json:"type"`
	Xml        map[string]string`json:"xml"`
	Properties map[string]interface{} `json:"properties"`
}

type Tags struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}
