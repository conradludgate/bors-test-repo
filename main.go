package main

import (
	"fmt"
)

func bifurcate() {
	fmt.Println("bifurcate called")
}

func main() {
	bifurcate()

	fmt.Println("hello world")

	bifurcateCrab()
}
