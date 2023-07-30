package main

import "fmt"

// 结构体的使用，类似与JavaBean
type Person struct {
	Name   string
	Age    int
	Scores [5]int
	ptr    *int
	slice  []int
	mapa   map[string]string
}

func main() {

	var p1 Person
	p1.slice = make([]int, 2)
	p1.slice[0] = 100
	p1.mapa = make(map[string]string)
	p1.mapa["a"] = "b"
	fmt.Println(p1)
}
