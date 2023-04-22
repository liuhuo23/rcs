package rcs

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Client struct {
	Host    string
	Port    uint
	Request Request
	conn    net.Conn
}

var err error

func ClientDefault() Client {
	return Client{
		Port:    8080,
		Host:    "127.0.0.1",
		Request: Request{},
		conn:    nil,
	}
}

func NewClient(host string, port uint) *Client {
	return &Client{
		Port: port,
		Host: host,
	}
}

func (c *Client) Set(value string) {
	c.Request.Method = CREATE
	c.Request.Time = time.Now()
	c.Request.Content = []byte(value)
}

func (c *Client) Del(value string) {
	c.Request.Method = DELETE
	c.Request.Content = []byte(value)
}

func (c *Client) Update(oldValue, newValue string) {
	c.Request.Method = UPDATE
	c.Request.Content = []byte(oldValue + " " + newValue)
}

func (c *Client) Commit() error {
	sendByte := c.Request.ToByte()
	length := fmt.Sprintf("%d\n", len(sendByte))
	sendByte = append([]byte(length), sendByte...)
	n, err := c.conn.Write(sendByte)
	fmt.Println("client n:", n)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetResult() string {
	reader := bufio.NewReader(c.conn)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "error"
	}
	return str
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) IsLive() bool {
	if err != nil {
		return false
	}
	return true
}

func (c *Client) Connect() error {
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("err : ", err)
		return err

	}
	c.conn = conn
	return nil
}
