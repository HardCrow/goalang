package main

import "fmt"

func main() {
	//map的定义
	//var 名字 map[key类型]value的类型
	var a map[string]string                                                //方式一
	var cities = make(map[string]string)                                   //方式二
	var city map[string]string = map[string]string{"01": "江西", "02": "北京"} //方式三
	//在使用map前要make，为了给map分配数据空间，和数组不一样，数组定义完后就直接分配空间
	a = make(map[string]string, 10) //map[string]string是类型，10是空间
	a["huang"] = "yingjie"
	a["1"] = "2"
	cities["2"] = "4"
	fmt.Println(cities)
	fmt.Println(city)
	fmt.Println(a) //运行后得出golang的map存储是无序的

	for k, v := range city {
		fmt.Printf("%v,%v", k, v) //for range遍历map
	}
}
