package hello

import "fmt"

// PrintHello public
func PrintHello() {
	fmt.Println("hello world")
}

// private
var aa = 12
var AA = 34

func Test() {

	fmt.Printf("局部param %d\n", aa)
	fmt.Printf("全局param %d\n", AA)
}
