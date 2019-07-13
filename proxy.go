package main

import (
	"log"
	"net/http"
	"websocket-dynamic-reverse-proxy/proxy"
)


func main() {
	err := http.ListenAndServe("0.0.0.0:8181", proxy.NewProxy())
	if err != nil {
		log.Fatalln(err)
	}
}
