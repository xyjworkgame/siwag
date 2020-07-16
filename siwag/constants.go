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
	License := spec.License{
		Name: "Apache 2.0",
		URL: "http://www.apache.org/licenses/LICENSE-2.0.html",
	}
	InfoBasic = spec.InfoProps{
		Description:    "you think this info",
		Version:        "1.0.0",
		Title:          "you think this title",
		TermsOfService: "",
		Contact:        &Contact,
		License: &License,
	}

	info := spec.Info{
		InfoProps: InfoBasic,

	}
	securitySchemeProps := spec.SecuritySchemeProps{}
	InitInfo = model.Base{
		Swagger:  "2.0",
		Info:     &info,
		Host:     "localhost:8081",
		BasePath: "/",
		//Tags:
		Schemes:  []string{"https",
			"http"},
		SecurityDefinitions:  &securitySchemeProps,
		//Definitions:
		//ExternalDocs: TODO 暂时不写
	}
}
