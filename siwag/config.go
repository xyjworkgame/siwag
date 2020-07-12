package siwag

/*
   add custom config
*/

type Config struct {

    IsOn bool
	// doc title
	DocTitle string
	// doc path
	DocPath string
	// 作者
	Author string
	// 邮件地址
	Email string
	// 描述
	Description string

	Host     string
	BasePath string
}
