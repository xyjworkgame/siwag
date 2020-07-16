package util

import (
	"log"
	"os"
)

func IsFileExist(path string)(bool){

	stat, err := os.Stat(path)
	if err != nil{
		log.Println(err)
		return false
	}
	if stat.IsDir(){
		return false
	}
	return true
}
