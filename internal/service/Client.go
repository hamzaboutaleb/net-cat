package service

import (
	"fmt"
	"net"
	"sync"

	"netcat/internal/utils"
)

type Client struct {
	Clients map[net.Conn]string
	Names   map[string]bool
	mut     sync.Mutex
}

func InitClient(c *Client) {
	c.Clients = make(map[net.Conn]string)
	c.Names = map[string]bool{}
}

func (c *Client) Add(conn net.Conn, name string) bool {
	if c.isUserExit(name) {
		return false
	}
	c.mut.Lock()
	c.Clients[conn] = name
	c.Names[name] = true
	c.mut.Unlock()
	return true
}

func (c *Client) CloseAll() {
	for key := range c.Clients {
		key.Close()
	}
}

func (c *Client) GetName(conn net.Conn) string {
	value := c.Clients[conn]
	return value
}

func (c *Client) isUserExit(name string) bool {
	c.mut.Lock()
	_, exist := c.Names[name]
	c.mut.Unlock()
	return exist
}

func (c *Client) BroadCastMessage(o any) {
	for conn := range c.Clients {
		message := fmt.Sprint(o)
		conn.Write([]byte(message))
	}
}

func (c *Client) BroadCastExcept(self net.Conn, o any) {
	for conn := range c.Clients {
		if conn != self {
			message := fmt.Sprintf("\n%v", o)
			conn.Write([]byte(message))
			utils.Prompt(conn, c.GetName(conn))
		}
	}
}

func (c *Client) delete(conn net.Conn) bool {
	name, ok := c.Clients[conn]

	if ok {
		delete(c.Clients, conn)
		delete(c.Names, name)
		return true
	}

	return false
}
