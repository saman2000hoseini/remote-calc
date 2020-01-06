package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"simpleCalc/model"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		var op1, op2 float64
		var operator string
		fmt.Fscan(input, &op1)
		fmt.Fscan(input, &operator)
		fmt.Fscan(input, &op2)
		op := &model.Operation{op1, op2, operator, 0}
		buf := bytes.NewBuffer(nil)
		json.NewEncoder(buf).Encode(op)
		response, err := http.Post("http://localhost:65431/calculate", "application/json; charset=utf-8", buf)
		if err != nil {
			panic(err)
		}
		json.NewDecoder(response.Body).Decode(&op)
		fmt.Println(op.Result)
	}
}
