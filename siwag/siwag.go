package siwag

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"yaagOrSwaggerDemo/siwag/models"
	"yaagOrSwaggerDemo/util"
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
		//log.Println(dataFile)
		//log.Println(InitInfo)
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

	// FIXME 尝试续上

	filePath, err := filepath.Abs(config.DocPath + ".json")
	if err != nil {
		log.Println(err)
		return
	}
	//	1. 首先判断，是否存在这个文件
	exist := util.IsFileExist(filePath)

	// 文件不存在
	if !exist {
		dataFile, err := os.Create(filePath)
		if err != nil {
			log.Println(err)
			return
		}
		//	2. 如果不存在，打开文件并且将json写入文件中
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
		return
	}
	//文件存在
	//	2. 如果存在文件，读取文件中已经存在的数据，
	initInfo := models.Base{}
	fileContent, err := os.OpenFile(filePath, os.O_RDWR, 6)
	fileData, err := ioutil.ReadAll(fileContent)
	if err != nil {
		log.Println(err)
		return
	}
	defer fileContent.Close()
	//	3. 如何去替换更新数据，然后再存储到里面去
	if err := json.Unmarshal([]byte(fileData), &initInfo); err != nil {
		log.Println(err)
		return
	}



	compareInitInfo(InitInfo, &initInfo)
	dataFile, err := os.Create(filePath)
	defer dataFile.Close()
	marshal, err := json.Marshal(InitInfo)
	if err != nil {
		log.Println(err)
		return
	}
	//fileContent.Truncate()
	// 移动指针
	//_, _ = fileContent.Seek(0, 0)
	//fileContent.Truncate(-1)
	//lower := strings.ToLower(string(marshal))
	if _, err = dataFile.Write(marshal); err != nil {
		log.Println(err)
		return
	}

	log.Println("finish record")
}

func compareInitInfo(nInfo *models.Base, oInfo *models.Base) {
	//oldInfo.Paths
	//oldInfo.Definitions
	//	 需要替换的上面两个
	// 以新的为主，旧的为参考

	for pathItemsk, pathItemsv := range oInfo.Paths {
		if _, ok := nInfo.Paths[pathItemsk]; ok {
			//	 path -url 存在
			for pathk,pathv := range pathItemsv{
				if _,ok := nInfo.Paths[pathItemsk][pathk];ok{
					// path -url - method 存在
					// compare parameter
					parameters := compareParameters(nInfo.Paths[pathItemsk][pathk].Parameters,pathv.Parameters)
					//nInfo.Paths[pathItemsk][pathk].Parameters = parameters
					log.Println(parameters)
					// FIXME 为了map值的正确，go语言不允许直接修改map中的值类型结构。
					// 怎么可以让这句话，添加一个伪遍历
					for k,_ := range nInfo.Paths{
						if k == pathItemsk{
							for _,v1 := range nInfo.Paths[pathItemsk]{
								v1.Parameters = parameters
							}
						}
					}
					// response
					for resk, resv := range oInfo.Paths[pathItemsk][pathk].Responses {
						//	2. 判断response
						if _, ok := pathv.Responses[resk]; ok {
						} else {
							//	不相同的则添加
							pathv.Responses[resk] = resv
						}
					}
				}else {
					nInfo.Paths[pathItemsk][pathk] = pathItemsv[pathk]
				}
			}
		} else {
			//	path-url 不存在，直接添加
			nInfo.Paths[pathItemsk] = oInfo.Paths[pathItemsk]
		}
	}
}






func compareParameters(parameters models.Parameters, oldP models.Parameters) []models.Parameter {
	var result []models.Parameter
	nameMap := map[string]models.Parameter{}
	for _, v := range parameters {
		if _, ok := nameMap[v.BodyParameter.Name]; !ok {
			//	name存在，
			//	result = append(result, v)
			nameMap[v.BodyParameter.Name] = v
		}

	}
	for _, v := range oldP {
		if _, ok := nameMap[v.BodyParameter.Name]; !ok {
			//	如果没有，则设置required
			v.BodyParameter.Required = false
			nameMap[v.BodyParameter.Name] = v
		}
	}
	for _, v := range nameMap {
		result = append(result, v)
	}

	// query
	for _, v := range parameters {
		if _, ok := nameMap[v.QueryParameter.Name]; !ok {
			//	name存在，
			//	result = append(result, v)
			nameMap[v.QueryParameter.Name] = v
		}

	}
	for _, v := range oldP {
		if _, ok := nameMap[v.QueryParameter.Name]; !ok {
			//	如果没有，则设置required
			v.QueryParameter.Required = false
			nameMap[v.QueryParameter.Name] = v
		}
	}
	for _, v := range nameMap {
		result = append(result, v)
	}
	return result
}
