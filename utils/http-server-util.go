package utils

import (
	"encoding/json"
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
	json.NewEncoder(w).Encode(op.Result)
}
