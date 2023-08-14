package main

import (
	"fmt"
	"net"
)

func Acc(con net.Conn) {
	defer con.Close()

	for {
		buf := make([]byte, 1024)
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("对方退出了对话", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("err为:", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待链接")

		accept, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("Accpet出错", err)
		} else {
			fmt.Println("成功", accept)
			fmt.Println("", accept.RemoteAddr().String())
		}
		go Acc(accept)
	}
	//fmt.Printf("success:%v\n",listen)
}
