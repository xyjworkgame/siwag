# 改写yaag,生成swagger.json文件



1. 目前只写出适用于gin框架的中间件
2. 用法如下：
```go
config := siwag.Config{
		IsOn:     true,
		DocTitle: "Gin",
		DocPath:  "apidoc",
		Author:   "demo",
		Email:    "1111@163.com",
		Host:     "172.0.0.1",
		BasePath: "/",
	}
	//add model
	siwag.AutoCreateJson(model.User{}, model.Permission{})
	//	init config file
	siwag.Init(&config)
``` 


3. 缺点： 
    1. 对于params 的参数，无法区别出来是什么数据，后端得到的数据都是字符串
    2. tag 标签还没有优化
    3. body 参数的type 无法很好的取消掉（没有很好的struct支持）
