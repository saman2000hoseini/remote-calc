package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"simpleCalc/model"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:65431")
	defer connection.Close()
	if err != nil {
		panic(err)
		return
	}
	input := bufio.NewReader(os.Stdin)
	client := model.NewClient(connection)
	go client.StartResponseHandler()
	for {
		var op1, op2 float64
		var operator string
		fmt.Fscan(input, &op1)
		fmt.Fscan(input, &operator)
		fmt.Fscan(input, &op2)
		op := &model.Operation{op1, op2, operator, 0}
		req, err := json.Marshal(op)
		if err != nil {
			panic(err)
		}
		client.Request(string(req))
	}
}
