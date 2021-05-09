package server

import (
	"context"
	"fmt"
	"io"
	"net"
)

type tcpServer struct {
	listener net.Listener
	addr     string
	ctx      context.Context
}

func handleRequest(conn net.Conn) {
	fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
	defer conn.Close()
	for {
		_, err := io.Copy(conn, conn)
		if err != nil {
			fmt.Printf("server: connect over  %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
			return
		}
	}
}

func (s *tcpServer) Start() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer l.Close()
	s.listener = l
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go handleRequest(conn)
	}
}

func (s *tcpServer) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}

func CreateTcpServer(addr string, ctx context.Context) *tcpServer {
	return &tcpServer{addr: addr, ctx: ctx}
}
