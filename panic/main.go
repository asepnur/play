package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	err := fmt.Errorf("panic here")
	defer func() {
		if r := recover(); err != nil {
			_ = r
			log.Println("trace: ", string(debug.Stack()))
		}
	}()
	doPanic(err)
}

func doPanic(err error) {
	panic(err)
}
