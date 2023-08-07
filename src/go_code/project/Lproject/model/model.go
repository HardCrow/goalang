package model

import "fmt"

// 封装性 小写代表私有（仅本包有用） 大写代表public  私有的话用工厂模式在外部调用函数即可
// 字段 小写则用func（方法取出）   大写可直接调用
type account struct {
	accountNo string
	pwd       string
	balance   float64
}

// 因为student小写了，外部看不见，利用工厂模式使外部看见，并不暴露内部信息
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度在6-10之间")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码必须是6位")
		return nil
	}
	if balance <= 20 {
		fmt.Println("存款太小")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}
func (a account) Set(acc string, pwd string, balance float64) *account {
	return &account{
		accountNo: acc,
		pwd:       pwd,
		balance:   balance,
	}
}

func (a account) Get() *account {
	return &account{
		accountNo: a.accountNo,
		pwd:       a.pwd,
		balance:   a.balance,
	}
}
