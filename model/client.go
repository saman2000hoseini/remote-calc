package model

import (
	"bufio"
	"encoding/json"
	"net"
)

type Client struct {
	connection net.Conn
	writer     *bufio.Writer
	reader     *bufio.Reader
	req        chan string
	connected  chan bool
}

func NewClient(connection net.Conn) *Client {
	return &Client{connection: connection, writer: bufio.NewWriter(connection), reader: bufio.NewReader(connection)}
}

func (c *Client) startClientHandler() {
	for {
		select {
		case req := <-c.req:
			op, err := NewOperation(req)
			if err != nil {
				panic(err)
				break
			}
			er := op.Calculate()
			if er != nil {
				panic(er)
				break
			}
			response, _ := json.Marshal(op)
			c.writer.WriteString(string(response))
		case connected := <-c.connected:
			if !connected {
				c.connection.Close()
				return
			}

		}
	}
}

func (c *Client) Listen(serverSide bool) {
	if serverSide {
		go c.startClientHandler()
	}
	for {
		req, err := c.reader.ReadString('\n')
		if err != nil {
			panic(err)
			c.connected <- false
		} else {
			c.req <- req
		}
	}
}
