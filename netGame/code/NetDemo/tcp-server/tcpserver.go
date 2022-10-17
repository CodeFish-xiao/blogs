package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8972")
	if err != nil {
		panic(err)
	}
	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()
	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Timeout() {
				log.Printf("accept temp err: %v", ne)
				continue
			}
			log.Printf("accept err: %v", e)
			return
		}
		go handleConn(conn)
		connections = append(connections, conn)
		if len(connections)%100 == 0 {
			log.Printf("total number of connections: %v", len(connections))
		}
	}
}

//处理链接
func handleConn(conn net.Conn) {
	c := NewClient(conn)
	go c.Echo()
}

// Client 客户端结构体
type Client struct {
	net.Conn //保存链接
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn}
}

// Echo Echo逻辑代码，收到消息直接发回客户端即可
func (c Client) Echo() {
	buf := make([]byte, 1024)
	for {
		c.Read(buf)
		c.Write(buf)
	}
}
