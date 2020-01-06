package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Operation struct {
	Operand1 float64 `json:"operand1"`
	Operand2 float64 `json:"operand2"`
	Operator string  `json:"operator"`
	Result   float64 `json:"result"`
}

func (o *Operation) Calculate() error {
	switch o.Operator {
	case "+":
		o.Result = o.Operand1 + o.Operand2
	case "-":
		o.Result = o.Operand1 - o.Operand2
	case "*":
		o.Result = o.Operand1 * o.Operand2
	case "/":
		o.Result = o.Operand1 / o.Operand2
	default:
		return errors.New("Something goes wrong")
	}
	return nil
}

func NewOperation(op string) (*Operation, error) {
	operation := &Operation{}
	if err := json.Unmarshal([]byte(op), operation); err != nil {
		return nil, err
	}
	fmt.Println(operation)
	return operation, nil
}
