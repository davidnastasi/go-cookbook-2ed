package main

import (
	"fmt"
	"net"
)

const addr = "localhost:8888"

func main() {

	conns := &connections{
		addrs: make(map[string]*net.UDPAddr),
	}

	fmt.Printf("serving on %s\n", addr)

	addr, err := net.ResolveUDPAddr("udp",addr)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp",addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go broadcast(conn, conns)

	msg := make([]byte,1024)
	for {
		// receive a message to gather the ip address and port to send back to
		_, retAddr, err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}

		conns.mu.Lock()
		conns.addrs[retAddr.String()] = retAddr
		conns.mu.Unlock()
		fmt.Printf("%s connected \n ", retAddr)
	}


}
