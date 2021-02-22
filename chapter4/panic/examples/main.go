package main

import (
	"fmt"
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
