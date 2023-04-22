package rcs

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	request := Request{
		Method:   CREATE,
		Send:     "echo hello",
		Time:     time.Now(),
		Sign:     "sdfsdfs",
		UserName: "sdfsdf",
		PassWord: "sdfsdf",
		Content:  []byte("sdfsdfsdfs"),
	}
	fmt.Println(request.ToByte())
}

func TestInt(t *testing.T) {
	str := "31"
	length, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(length)
}
