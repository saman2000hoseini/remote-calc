package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simpleCalc/model"
)

func ServerHandler(w http.ResponseWriter, req *http.Request) {
	op := new(model.Operation)
	err := json.NewDecoder(req.Body).Decode(&op)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	err = op.Calculate()
	if err != nil {
		panic(err)
	}
	fmt.Println(op.Result)
	json.NewEncoder(w).Encode(op)
}
