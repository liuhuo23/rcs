package rcs

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	host  string
	port  uint
	debug bool
}

func NewServer(host string, port uint, debug bool) Server {
	app := Server{
		host:  host,
		port:  port,
		debug: debug,
	}
	return app
}

func ServerDefault() Server {
	app := Server{
		host:  "127.0.0.1",
		port:  8080,
		debug: true,
	}
	return app
}

func (a Server) Run() {
	fmt.Println("服务端启动。。。。")
	address := fmt.Sprintf("%s:%d", a.host, a.port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listen() failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() failed, err:", err)
			continue
		}
		fmt.Println(conn.RemoteAddr().String())
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	fmt.Println(conn.RemoteAddr())
	//var buf [1024]byte
	for {
		//n, _ := reader.Read(buf)
		////fmt.Println("N:", n)
		//if n == 0 {
		//	fmt.Println(buf[:n])
		//	continue
		//}
		//request := ParseRequest(string(buf[:n]))
		//if request == nil {
		//	fmt.Print("request:", request)
		//	conn.Write([]byte("error"))
		//	continue
		//}
		//fmt.Printf("request: %v\n", request)
		//conn.Write([]byte("ok"))
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from conn failed, err:%v", err)
			return
		}
		length, err := strconv.ParseInt(strings.Trim(str, "\n"), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("length", length)
		rcvByte, err := reader.Peek(int(length))
		if err != nil {
			return
		}
		fmt.Println(string(rcvByte))
		reader.Reset(conn)
		_, err = conn.Write([]byte("ok\n"))
		if err != nil {
			continue
		}
	}
}
