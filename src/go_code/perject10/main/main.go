package main

import "fmt"

func main() {
	//切片的使用方法二
	//var 切片名 【】type =make（【】type，len，【cap】）
	var slice []float64 = make([]float64, 5, 10)
	slice[0] = 20
	slice[3] = 60
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	//方式三 切片
	var slice1 []string = []string{"tom", "jack"}
	fmt.Println(slice1)
	//下面是数组
	var arr [3]int = [3]int{1, 2, 3}
	slice3 := arr[0:2]
	fmt.Println(slice3)
	fmt.Println(arr)
	var slice4 []int = []int{5, 6, 5, 52, 4, 3}
	fmt.Println(slice4)
	//切片追加 加属性
	slice5 := append(slice4, 52, 90)
	fmt.Println(slice5)
}
