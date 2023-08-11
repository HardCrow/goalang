package main

import (
	"encoding/json"
	"fmt"
)

type Moster struct {
	Name string //'json:"name"'//反序列后的名字  利用反射实现
	Age  int
}

// 反序列化  用json反序列化成struct
func unmarshal() {
	str := "{\"Name\":\"A\",\"Age\":12}"
	var mos Moster
	err := json.Unmarshal([]byte(str), &mos)
	if err != nil {
		fmt.Println("反序列失败")
	}
	fmt.Printf("%v\n", mos)
}
func main() {
	unmarshal()
}
