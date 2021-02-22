package main

import (
	"fmt"
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter4/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
