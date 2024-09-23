package service

import (
	"fmt"
	"net"

	"netcat/internal/config"
	"netcat/internal/utils"
)

const (
	INCREMENT = iota
	DECREMENT
)

type Server struct {
	Port        string
	Protocol    string
	Listener    net.Listener
	ClientCount int
	Locker      MutexLock
}

func NewServer() *Server {
	return &Server{
		Port:        ":8000",
		Protocol:    "tcp",
		ClientCount: 0,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen(s.Protocol, s.Port)
	if err != nil {
		return err
	}
	s.Listener = listener
	return nil
}

func (s *Server) Accept() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		s.HandleConnection(conn)
	}
}

func (s *Server) SetPort(port string) {
	s.Port = ":" + port
}

func (s *Server) ClientCountLock() {
	s.Locker.ClientCount.Lock()
}

func (s *Server) ClientCountUnlock() {
	s.Locker.ClientCount.Unlock()
}

func (s *Server) IsServerFull() bool {
	s.ClientCountLock()
	if s.ClientCount >= config.MAX_CLIENT {
		s.ClientCountUnlock()
		return false
	}
	s.ClientCountUnlock()
	return true
}

func (s *Server) ChangeClientCount(state int) {
	s.ClientCountLock()
	switch state {
	case INCREMENT:
		s.ClientCount++
	case DECREMENT:
		s.ClientCount--
	}
	s.ClientCountUnlock()
}

func (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()
	if s.IsServerFull() {
		fmt.Println(config.ErrClientFull)
		conn.(*net.TCPConn).SetLinger(0)
		return
	}
	s.ChangeClientCount(INCREMENT)
	utils.PrintLogo(conn)
}
