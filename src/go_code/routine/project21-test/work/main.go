package work

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Restore() bool {
	filePath := "D:/GoLangCode/src/go_code/mos.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err2 := json.Unmarshal(data, this)
	if err2 != nil {
		fmt.Println(err2)
		return false
	}
	return true
}
func (this *Monster) Store() bool {

	marshal, err := json.Marshal(this)
	if err != nil {
		fmt.Println("err:", err)
		return false
	}
	filePath := "D:/GoLangCode/src/go_code/mos.ser"
	err1 := ioutil.WriteFile(filePath, marshal, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return false
	}
	return true
}
