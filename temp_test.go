package main

import (
	"testing"
	"yaagOrSwaggerDemo/model"
)

/*
@Time : 2020/7/12 21:52
@Author : Firewine
@File : temp_test
@Software: GoLand
@Description:
*/

func AutoCreateJson(values ...interface{}){
	// 直接重新覆盖原先的数据
	print(values[0])

}
func Test(t *testing.T) {
	AutoCreateJson(&model.User{},model.Permission{})

}
