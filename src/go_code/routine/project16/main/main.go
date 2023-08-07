package main

import (
	"fmt"
	"go_code/routine/project16/model"
)

func main() {
	var stu = model.Student{
		Name:  "tom",
		Score: 70,
	}
	var tea = model.NewTeacher("a", 90)
	fmt.Println(stu)
	fmt.Println(*tea)
}
