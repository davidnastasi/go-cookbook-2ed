package main

import (
	"fmt"
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter4/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()


}