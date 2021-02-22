package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8888"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter some text: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading input: %s\n", err.Error())
			continue
		}
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("encountered an error connecting: %s\n", err.Error())
		}
		fmt.Fprintf(conn, data)
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading response: %s\n", err.Error())
		}
		fmt.Printf("Received back: %s", status)
		conn.Close()
	}

}