package main

import (
	"fmt"
	"github.com/go-programing-cookbook-2ed/chapter4/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()


}