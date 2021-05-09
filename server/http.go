package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func connectCheck(conn net.Conn, state http.ConnState) {
	if state == http.StateNew {
		tcpConn, ok := conn.(*net.TCPConn)
		if ok {
			tcpConn.SetLinger(30)
		}
		fmt.Printf("新的客户端连接 = %v\n", conn.RemoteAddr())
	}

	if state == http.StateClosed {
		fmt.Printf("客户端断开连接 = %v\n", conn.RemoteAddr())
	}
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

type httpServer struct {
	server *http.Server
	ctx    context.Context
}

func (s *httpServer) init() error {
	http.HandleFunc("/", ServeHome)
	return nil
}

func (s *httpServer) Start() error {
	s.init()
	return s.server.ListenAndServe()
}

func (s *httpServer) Stop() error {
	s.server.Shutdown(s.ctx)
	s.server.Close()
	return nil
}

func CreateHttpServer(addr string, ctx context.Context) *httpServer {
	s := &http.Server{
		ConnState:         connectCheck,
		Addr:              addr,
		ReadTimeout:       20 * time.Second,
		ReadHeaderTimeout: 20 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       90 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	return &httpServer{server: s, ctx: ctx}
}
