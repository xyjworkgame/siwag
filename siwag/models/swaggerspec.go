package models

import "github.com/go-openapi/spec"

// 这里搭建json 的框架
type SwaggerSpec struct {
	Base  spec.SwaggerProps `json:"base"`
	Paths Paths             `json:"paths"`
}
