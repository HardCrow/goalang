package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.177.1:8888")
	if err != nil {
		fmt.Println("失败", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	//从终端读取一行输入并发给服务器
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("用户数据读取失败")
	}
	//再将readString发给服务器
	n, err := conn.Write([]byte(readString))
	if err != nil {
		fmt.Println("conn.write失败", err)
	}
	fmt.Println("客户端发送了%d字节并退出", n)
}
