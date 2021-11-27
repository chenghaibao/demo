package utils

import (
	"fmt"
	"os"
)

func CheckFile(Filename string) bool {
	var exist = true
	if _, err := os.Stat(Filename); os.IsNotExist(err) {
		exist = false
		if err != nil{
			fmt.Println("检查文件失败")
		}
	}
	return exist
}