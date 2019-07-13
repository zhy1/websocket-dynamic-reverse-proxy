package main

import (
	"log"
	"net/http"
	"websocket-dynamic-reverse-proxy/custom"
)


func main() {
	err := http.ListenAndServe("0.0.0.0:8181", custom.NewProxy())
	if err != nil {
		log.Fatalln(err)
	}
}
