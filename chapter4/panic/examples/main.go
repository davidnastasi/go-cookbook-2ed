package main

import (
	"fmt"
	"github.com/davidnastasi/go-programing-cookbook-2ed/chapter4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
