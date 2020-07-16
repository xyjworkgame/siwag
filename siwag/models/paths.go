package models

import (
	"github.com/go-openapi/spec"
)

// Paths 是一个字典，里面是每一个json
// 尝试查看调用方法，反射url路径
type Paths map[string]PathItems
type PathItems map[string]Path
type BodyParameters []BodyParameter
type QueryParameters []BodyParameter
type Parameters []Parameter

// Path 是一个字典，url -》 {}
type Path struct {
	Description string   `json:"description"`
	Consumes    []string `json:"consumes"`
	Produces    []string `json:"produces"`
	Schemes     []string `json:"schemes"`
	Tags        []string `json:"tags"`
	Summary     string   `json:"summary"`
	//ID          string           `json:"id"`
	Deprecated bool             `json:"deprecated"`
	Parameters Parameters       `json:"parameters"`
	Responses  map[int]Response `json:"responses"`
}

type QueryParameter struct {
	In              string `json:"in"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Required        bool   `json:"required"`
	AllowEmptyValue bool   `json:"allowEmptyValue"`
	Description     string `json:"description"`
}
type BodyParameter struct {
	In              string `json:"in"`
	Name            string `json:"name"`
	Required        bool   `json:"required"`
	AllowEmptyValue bool   `json:"allowEmptyValue"`
	Description     string `json:"description"`
}
type Parameter struct {
	In              string `json:"in"`
	Name            string `json:"name"`
	Required        bool   `json:"required"`
	Type            string `json:"type"`
	AllowEmptyValue bool   `json:"allowEmptyValue"`
	Description     string `json:"description"`
}
type Schema struct {
	Type        string      `json:"type"`
	Item        *spec.Items `json:"item"`
	Description string      `json:"description"`
}

type Response struct {
	Description string `json:"description"`
}
