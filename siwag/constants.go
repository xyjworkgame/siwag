package siwag

import (
	"github.com/go-openapi/spec"
	model "yaagOrSwaggerDemo/siwag/models"
)

// Initial empty spec
var InitInfo  model.Base
var Contact spec.ContactInfo
var InfoBasic spec.InfoProps

// Initial Basics spec
func init() {

	Contact = spec.ContactInfo{
		Name:  "",
		Email: "",
		URL:   "",
	}
	InfoBasic = spec.InfoProps{
		Description:    "you think this info",
		Version:        "1.0.0",
		Title:          "you think this title",
		TermsOfService: "",
		Contact:        &Contact,
	}
	info := spec.Info{
		InfoProps: InfoBasic,
	}
	InitInfo = model.Base{
		Swagger:  "2.0",
		Info:     &info,
		Host:     "localhost:8081",
		BasePath: "/",
		//Tags: TODO 1. 自动根据path抽取，或者手动添加
		//Schemes: spec.Schema{}  TODO 暂时不写
		//SecurityDefinitions:  TODO 暂时不写
		//Definitions:
		//ExternalDocs: TODO 暂时不写
	}
}
