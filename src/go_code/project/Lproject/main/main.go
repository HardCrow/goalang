package main

import (
	"fmt"
	"go_code/project/Lproject/model"
)

func main() {

	account := model.NewAccount("12345687", "123456", 60)
	if account != nil {
		fmt.Println("创建成功")
		fmt.Println(*account)
	} else {
		fmt.Println("重新输入")

	}

}
