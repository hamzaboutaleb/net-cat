package service

import (
	"bufio"
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
	server.History.Logger = InitLogger()
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
		fmt.Fprintln(conn, config.ErrClientFull)
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
	if s.Clients.Delete(conn) {
		s.ChangeClientCount(DECREMENT)
	}
}

func (s *Server) PromptName(conn net.Conn) string {
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
	return username
}

func (s *Server) PromptChat(conn net.Conn, username string) {
	scanner := bufio.NewScanner(conn)
	utils.Prompt(conn, username)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 && len(strings.TrimSpace(text)) > 0 && utils.IsPrint(text) {
			text += "\n"
			message := Message{
				Date:     time.Now(),
				Msg:      text,
				Username: s.Clients.GetName(conn),
			}
			s.History.Push(message)
			s.Clients.BroadCastExcept(conn, message)
			fmt.Println("error")
		}
		utils.Prompt(conn, username)
	}
}

func (s *Server) HandleConnection(conn net.Conn) {
	defer conn.Close()
	defer s.ClientLogout(conn)
	defer fmt.Printf("number of connetced client %v -- %v\n", s.ClientCount, s.Clients.Clients)

	if !s.CheckNewClient(conn) {
		return
	}
	utils.PrintLogo(conn)
	username := s.PromptName(conn)
	s.PromptChat(conn, username)
}
