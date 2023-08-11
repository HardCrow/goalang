package main

import (
	//转成json的包
	"encoding/json"
	"fmt"
)

// 序列号操作  序列化成json串容易传输
type Moster struct {
	Name string //'json:"name"'//反序列后的名字  利用反射实现
	Age  int
}

func testSt() {
	moster := Moster{ //实例化
		Name: "A",
		Age:  12,
	}
	data, err := json.Marshal(&moster) //转成json字符串赋值给data 且判断是否有错误
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("序列化后", string(data))
}
func main() {
	testSt()
}
