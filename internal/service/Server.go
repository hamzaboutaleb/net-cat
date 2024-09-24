package service

import (
	"fmt"
	"net"
	"strings"
	"time"

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
	Clients     Client
	History     History
}

func NewServer() *Server {
	server := &Server{
		Protocol:    "tcp",
		ClientCount: 0,
	}
	InitClient(&server.Clients)
	server.SetPort(config.PORT)
	return server
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
		go s.HandleConnection(conn)
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
	return s.ClientCount >= config.MAX_CLIENT
}

func (s *Server) ChangeClientCount(state int) {
	switch state {
	case INCREMENT:
		s.ClientCount++
	case DECREMENT:
		s.ClientCount--
	}
}

func (s *Server) BroadCastMessage(o any) {
	s.History.Push(o)
	s.Clients.BroadCastMessage(o)
}

func (s *Server) BroadCastHistory(c net.Conn) {
	s.History.PrintHistory(c)
}

func (s *Server) CheckNewClient(conn net.Conn) bool {
	s.ClientCountLock()
	defer s.ClientCountUnlock()
	if s.IsServerFull() {
		fmt.Println(config.ErrClientFull)
		conn.(*net.TCPConn).SetLinger(0)
		return false
	}
	s.ChangeClientCount(INCREMENT)
	return true
}

func (s *Server) ClientLogout(conn net.Conn) {
	s.ClientCountLock()
	defer s.ClientCountUnlock()
	name := s.Clients.GetName(conn)
	s.History.Push(utils.LeftMsg(name))
	s.Clients.BroadCastExcept(conn, utils.LeftMsg(name))
	if s.Clients.delete(conn) {
		s.ChangeClientCount(DECREMENT)
	}
}

func (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()
	defer s.ClientLogout(conn)
	defer fmt.Printf("number of connetced client %v -- %v\n", s.ClientCount, s.Clients.Clients)
	fmt.Printf("%v Connected\v", conn)
	if !s.CheckNewClient(conn) {
		return
	}
	utils.PrintLogo(conn)
	username := ""
	for {
		name, err := utils.GetName(conn)
		if err != nil {
			fmt.Fprintln(conn, err)
			conn.Write(utils.ToBytes("[ENTER YOUR NAME]:"))
			continue
		}
		if s.Clients.Add(conn, name) {
			s.BroadCastHistory(conn)
			s.History.Push(utils.JoinMsg(name))
			s.Clients.BroadCastExcept(conn, utils.JoinMsg(name))
			username = name
			break
		} else {
			fmt.Fprintln(conn, "name already exists")
			conn.Write(utils.ToBytes("[ENTER YOUR NAME]:"))
		}
	}

	// message := ""
	buffer := make([]byte, 1024*4)
	utils.Prompt(conn, username)
	for {
		n, err := conn.Read(buffer)

		fmt.Println(n, err)
		if n == 0 {
			return
		}
		// fmt.Println(n, err, string(buffer[:n-1]), buffer[n-1], strings.TrimSpace(string(buffer)[:n-1]))
		text := string(buffer)[:n]
		if text[n-1] != '\n' {
			text += "\n"
			conn.Write(utils.ToBytes("\n"))
		}
		if len(text) > 0 && len(strings.TrimSpace(text)) > 0 {
			message := Message{
				Date:     time.Now(),
				Msg:      text,
				Username: s.Clients.GetName(conn),
			}
			s.History.Push(message)
			s.Clients.BroadCastExcept(conn, message)
		}
		utils.Prompt(conn, username)

		// flushStdin(conn)
	}
}
