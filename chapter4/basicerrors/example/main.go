package main

import (
	"fmt"
	"github.com/davidnastasi/go-programing-cookbook-2ed/chapter4/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
