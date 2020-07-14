package models

import "github.com/go-openapi/spec"

// 这里搭建json 的框架
type SwaggerSpec struct {
	Base  spec.SwaggerProps // 基本信息
	paths Paths             // 路径所有信息
}
