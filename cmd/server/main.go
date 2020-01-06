package main

import (
	"net"
	"simpleCalc/model"
)

func main() {
	server, err := net.Listen("tcp", ":65431")
	if err != nil {
		panic(err)
	}
	for {
		connection, err := server.Accept()
		if err != nil {
			panic(err)
		}
		client := model.NewClient(connection)
		go client.Listen()
	}
}
