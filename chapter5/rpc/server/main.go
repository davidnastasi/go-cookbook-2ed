package main

import (
	"fmt"
	"github.com/go-programing-cookbook-2ed/chapter5/rpc/tweak"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	s := new(tweak.StringTweaker)
	if err := rpc.Register(s); err != nil {
		log.Fatal("failed to register:", err)
	}
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	fmt.Println("listen on :1234")
	log.Panic(http.Serve(l, nil))

}

