package rcs

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	Method   string
	Send     string
	Time     time.Time
	Sign     string
	UserName string
	PassWord string
	Content  []byte
}

func (r *Request) ToByte() []byte {
	var buf []byte
	method := fmt.Sprintf("%s:%s\n", METHOD, r.Method)
	buf = append(buf, []byte(method)...)
	if r.Send != "" {
		send := fmt.Sprintf("%s:%s\n", SEND, r.Send)
		buf = append(buf, []byte(send)...)
	}
	if r.UserName != "" {
		username := fmt.Sprintf("%s:%s\n", USERNAME, r.UserName)
		buf = append(buf, []byte(username)...)
	}
	if r.Sign != "" {
		sign := fmt.Sprintf("%s:%s\n", SIGN, r.Sign)
		buf = append(buf, []byte(sign)...)
	}
	Time := fmt.Sprintf("%s:%d\n", TIME, r.Time.Unix())
	buf = append(buf, []byte(Time)...)
	if r.PassWord != "" {
		password := fmt.Sprintf("%s:%s\n", PASSWORD, r.PassWord)
		buf = append(buf, []byte(password)...)
	}
	buf = append(buf, []byte("\n")...)
	buf = append(buf, r.Content...)
	return buf
}

// ParseRequest 解析请求体
func ParseRequest(receiveStr string) *Request {
	fmt.Println("开始")
	lines := strings.Split(receiveStr, "\n")
	print(len(lines))
	fmt.Println(receiveStr)
	request := &Request{}
	for _, line := range lines {
		keyValue := strings.Split(line, ":")
		if len(line) != 2 {
			continue
		}
		key := strings.Trim(keyValue[0], " ")
		value := strings.Trim(keyValue[1], " ")
		if strings.ToLower(key) == "method" {
			request.Method = strings.ToLower(value)
		} else if strings.ToLower(key) == "send" {
			request.Send = value
		} else if strings.ToLower(key) == "time" {
			timestamp, _ := strconv.ParseInt(value, 10, 64)
			request.Time = time.Unix(timestamp, 0)
		} else if strings.ToLower(key) == "sign" {
			request.Sign = value
		} else if strings.ToLower(key) == "username" {
			request.UserName = value
		} else if strings.ToLower(key) == "password" {
			request.PassWord = value
		}
	}
	return request
}
