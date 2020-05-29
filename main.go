package main

import "fmt"

func main() {
	var (
		a = 1
		b = 4
	)
	c := a + b
	fmt.Println("Hello, WebAssembly!")
	fmt.Println("a+b=", c)
	for i := 0; i < c; i++ {
		fmt.Println(i)
	}
}
