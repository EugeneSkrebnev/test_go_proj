package main

import (
	"fmt"
	"github.com/EugeneSkrebnev/test_go_lib/pkg/lib"
)

func main() {
	test := lib.NewLibObject("Hello, World!")
	fmt.Println(test.Path())
}
