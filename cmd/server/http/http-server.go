package main

import (
	"log"
	"net/http"
	"simpleCalc/utils"
)

func main() {
	http.HandleFunc("/calculate", utils.ServerHandler)
	log.Fatal(http.ListenAndServe(":65431", nil))
}
