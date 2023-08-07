package main

import "fmt"

// 三个班 每个班5个人
// 计算平均分和每个班总分
func main() {
	var student [3][5]int               //定义每个班的学生成绩
	for i := 0; i < len(student); i++ { //定义每个班的学生成绩
		for j := 0; j < len(student[i]); j++ { //定义每个班的学生成绩
			fmt.Println("请输入第%d班的第%d个学生的成绩\n", i+1, j+1) //定义每个班的学生成绩
			fmt.Scanln(&student[i][j])                   //定义每个班的学生成绩
		}
	}
	fmt.Println(student)
	for i := 0; i < len(student); i++ {
		sum := 0
		for j := 0; j < len(student[i]); j++ {
			sum += student[i][j]
		}
		fmt.Println(sum, sum/len(student[i]))
	}
}
