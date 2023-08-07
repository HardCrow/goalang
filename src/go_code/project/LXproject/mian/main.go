package main

import "fmt"

func main() {
	key := ""
	loop := true
	balance := 100.0
	money := 0.0
	note := ""
	detail := "收支\t账户金额\t收支金额\t说明\t"
	for {
		fmt.Println("-----------------家庭记账软件----------------")
		fmt.Println("1 收支明细")
		fmt.Println("2 登记支出")
		fmt.Println("3 登记支出")
		fmt.Println("4 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println(detail)
		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入\t%v\t\t\t%v\t%v\t\n", balance, money, note)
		case "3":
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if balance < money {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出\t%v\t\t\t%v\t%v\t\n", balance, money, note)
		case "4":
			loop = false
		default:
			fmt.Println("请输入正确选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出了使用")
}
