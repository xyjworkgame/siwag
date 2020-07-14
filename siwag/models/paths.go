package models

import "github.com/go-openapi/spec"

// Paths 是一个字典，里面是每一个json
// 尝试查看调用方法，反射url路径
type Paths map[string]Path

// Path 是一个字典，url -》 {}
type Path struct {
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
}

type Parameter struct {
	Description     string
	Name            string
	In              string
	Required        bool
	Schema          *spec.Schema //关于spec 全部 手动输入，
	AllowEmptyValue bool
}

type Responses struct {
	Default            *Response
	StatusCodeResponse map[int]Response
}

type Response struct {
	Description string
	Schema      *spec.Schema
	Headers     map[string]string
}
