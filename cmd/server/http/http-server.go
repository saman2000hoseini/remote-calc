package main

import (
	"log"
	"net/http"
	"simpleCalc/utils"
)

func main() {
	http.HandleFunc("/", utils.ServerHandler)
	log.Fatal(http.ListenAndServe("65432", nil))
}
