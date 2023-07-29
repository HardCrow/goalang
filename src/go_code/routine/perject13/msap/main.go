package main

import "fmt"

func main() {
	//map的切片操作
	var moters []map[string]string
	moters = make([]map[string]string, 2)
	if moters[0] == nil {
		moters[0] = make(map[string]string, 2)
		moters[0]["name"] = "xiaoxiao"
		moters[0]["age"] = "10"

	}
	fmt.Println(moters)
	a := make(map[int]int, 9)
	a[0] = 1
	a[8] = 2
	a[9] = 20
	a[5] = 110
	fmt.Println(a)
	//判断map是引用传递和值传递的方式
	map1 := make(map[int]int)
	map1[1] = 10
	map1[2] = 20
	map1[10] = 30
	modify(map1)
	fmt.Println(map1)
}
func modify(map1 map[int]int) {
	map1[10] = 900
}
