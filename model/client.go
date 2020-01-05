package model

import (
	"bufio"
	"net"
)

type Client struct {
	connection net.Conn
	writer *bufio.Writer
	reader *bufio.Reader
	req chan string
	connected chan bool
}
