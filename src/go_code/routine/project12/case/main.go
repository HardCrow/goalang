package main

import "fmt"

func main() {
	var student map[string]string = map[string]string{"小明": "男", "小红": "女", "大黄": "男"} //方法一
	studentMap := make(map[string]map[string]string)                                   //方法二
	studentMap["男"] = make(map[string]string, 3)
	studentMap["男"]["名字"] = "大荒"
	studentMap["男"]["名字"] = "大加" //覆盖map
	fmt.Println(studentMap)
	fmt.Println(student)
	val, ok := studentMap["男"] //map的查找
	if ok {
		fmt.Println("查到了", val)
	} else {
		fmt.Println("没查到")
	}
	for k, v := range studentMap {
		for k2, v2 := range v {
			fmt.Printf("%v,%v,%v", k, k2, v2) //for range map 遍历
		}
	}
	fmt.Println(len(student))
}
