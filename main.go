package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/techwo/rpc-server/calculator"
)

func main() {
	s := new(calculator.Service)
	rpc.Register(s)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "localhost:8767")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Println("Server started!")
	log.Println(l.Addr())
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
