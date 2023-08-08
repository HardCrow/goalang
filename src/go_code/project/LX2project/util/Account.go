package util

import "fmt"

type Account struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	detail  string
}

func NewAccount() *Account {
	return &Account{
		"",
		true,
		100.0,
		0.0,
		"",
		"收支\t账户金额\t收支金额\t说明\t",
	}
}
func (this *Account) showDetail() {
	fmt.Println(this.detail)
}
func (this *Account) incoming() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n收入\t%v\t\t\t%v\t%v\t\n", this.balance, this.money, this.note)
}
func (this *Account) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.balance < this.money {
		fmt.Println("余额不足")

	}
	this.balance -= this.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.detail += fmt.Sprintf("\n支出\t%v\t\t\t%v\t%v\t\n", this.balance, this.money, this.note)
}

//	func (this *Account) exit()  {
//		this.loop = false
//
// default:
//
//		fmt.Println("请输入正确选项")
//	}
func (this *Account) Select() {
	for {
		fmt.Println("-----------------家庭记账软件----------------")
		fmt.Println("1 收支明细")
		fmt.Println("2 登记支出")
		fmt.Println("3 登记支出")
		fmt.Println("4 退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetail()
		case "2":
			this.incoming()
		case "3":
			this.pay()
		case "4":
			this.loop = false
		default:
			fmt.Println("请输入正确选项")
		}
		if !this.loop {
			break
		}
	}

}
