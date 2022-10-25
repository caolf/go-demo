package main

import (
	"go-demo/array"
	"go-demo/goroutine"
	_interface "go-demo/interface"
)

func main() {

	array.TestArray()

	_interface.TestInterface()

	goroutine.TestGoroutine()
}
