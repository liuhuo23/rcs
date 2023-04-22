package main

import (
	"bufio"
	"fmt"
	"liuhuo23/rcs"
	"os"
	"strings"
)

var inputReader *bufio.Reader
var input string
var err error

// TCP 客户端
func main() {
	client := rcs.ClientDefault()
	client.Connect()
	defer client.Close()
	var prefix = fmt.Sprintf("%s:%d>", client.Host, client.Port)
	for {
		if !client.IsLive() {
			fmt.Println("连接已失效")
			client.Connect()
		}
		fmt.Printf("%s", prefix)
		inputReader = bufio.NewReader(os.Stdin)
		input, _ = inputReader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		if input == "Q" {
			fmt.Println("退出应用")
			break
		}
		client.Set(strings.Trim(input, "\r\n"))
		err := client.Commit()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(client.GetResult())
	}
}
