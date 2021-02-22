package main

import (
	"fmt"
	"github.com/davidnastasi/go-programing-cookbook-2ed/chapter5/dns"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address> \n", os.Args[0])
		os.Exit(1)
	}
	address := os.Args[1]
	lookup, err := dns.LookupAddress(address)
	if err != nil {
		log.Panic("failed to loopup: %s", err.Error())
	}
	fmt.Println(lookup)

}