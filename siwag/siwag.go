package siwag

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"yaagOrSwaggerDemo/siwag/models"
)

var config *Config

func IsOn() bool {
	return config.IsOn
}

// add config
// example : models
// example : title
func Init(conf *Config) {
	config = conf
	if conf.DocPath == "" {
		conf.DocPath = "apiSwagger"
	}
	// add custom info
	InitInfo.Host = config.Host
	InitInfo.BasePath = config.BasePath
	Contact.Email = config.Email
	Contact.Name = config.Author
	InfoBasic.Title = config.DocTitle
	InfoBasic.Description = config.Description
	// create file json
	//filePath, err := filepath.Abs(conf.DocPath + ".json")
	filePath, err := filepath.Abs(conf.DocPath + ".json")
	dataFile, err := os.Open(filePath)
	defer dataFile.Close()
	if err == nil {
		log.Println(dataFile)
		log.Println(InitInfo)
		json.NewDecoder(io.Reader(dataFile)).Decode(InitInfo)
		//generateHtml()
	}

}

func IsStatusCodeValid(code int) bool {
	if code >= 200 && code < 500 {
		return true
	} else {
		return false
	}
}

// reflect scan model
func AutoCreateJson(values ...interface{}) {
	definitions := make(map[string]*models.Definitions)
	for _, value := range values {

		refValue := reflect.ValueOf(value) // value
		refType := reflect.TypeOf(value)   // type
		fieldCount := refValue.NumField()  // field count
		//fmt.Println("fieldCount:", fieldCount)
		structName := refType.Name() // struct name

		definition := models.Definitions{}

		definition.Type = "object"

		properties := make(map[string]interface{})
		//fmt.Println("field name:", fieldName)
		for i := 0; i < fieldCount; i++ {
			fieldType := refType.Field(i) // field type
			//fmt.Println("field type:", fieldType.Type)
			//fmt.Println("field name1:", fieldType.Name)
			properties[fieldType.Name] = map[string]string{
				"type": fieldType.Type.String(),
			}

		}
		definition.Properties = properties
		definition.Xml = map[string]string{
			"name": structName,
		}
		definitions[structName] = &definition
	}
	InitInfo.Definitions = definitions
}

// 生成json
func GenerateJson(InitInfo *models.Base) {

	//	TODO 如何续传生成json文件


	//	3. 如何去替换更新数据，然后再存储到里面去
	// FIXME 尝试续上
	//	1. 首先判断，是否存在这个文件
	filePath, err := filepath.Abs(config.DocPath)
	if _, err := os.Stat(filePath); err != nil {
		log.Println(err)
		return
	} else {
		dataFile, err := os.Create(filePath + ".json")
		if err != nil {
			log.Println(err)
			return
		}
	}

	//	2. 如果存在文件，读取文件中已经存在的数据，
	FileData, err := ioutil.ReadAll(dataFile)
	json.Unmarshal()
	fmt.Println(dataFile)
	marshal, err := json.Marshal(InitInfo)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(marshal)
	defer dataFile.Close()
	_, err = dataFile.Write(marshal)
	if err != nil {
		log.Println(err)
		return
	}
}
