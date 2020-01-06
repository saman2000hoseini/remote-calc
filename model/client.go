package model

import (
	"bufio"
	"encoding/json"
	"fmt"
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

/*
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
			response, _ := json.Marshal(op.Result)
			fmt.Println(op.Result)
			c.writer.WriteString(string(response) + "\n")
			c.writer.Flush()
		case connected := <-c.connected:
			if !connected {
				c.connection.Close()
				return
			}

		}
	}
}
*/
func (c *Client) StartResponseHandler() {
	for {
		res, err := c.reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println("result = " + res)
	}
}

func (c *Client) Request(req string) {
	c.writer.WriteString(req + "\n")
	c.writer.Flush()
}

func (c *Client) Listen() {
	//	go c.startClientHandler()
	for {
		req, err := c.reader.ReadString('\n')
		if err != nil {
			panic(err)
			c.connected <- false
		} else {
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
			response, _ := json.Marshal(op.Result)
			fmt.Println(op.Result)
			c.writer.WriteString(string(response) + "\n")
			c.writer.Flush()

			//c.req <- req
		}
	}
}
