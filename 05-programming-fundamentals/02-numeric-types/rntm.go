package main

import (
	"fmt"
	"runtime"
)

func archTest() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
