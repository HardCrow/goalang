package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/GoLangCode/src/go_code/routine/testTXT")
	if err != nil {
		fmt.Printf("err", err)
	}
	fmt.Printf("file", file)
	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件错误")
	}
}
