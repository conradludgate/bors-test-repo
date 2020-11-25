package main

import "fmt"

func bifurcateCrab() {
	fmt.Println("bifurcateCrab called")
}

func main() {
	bifurcateCrab()

	fmt.Println("hello world")

	bifurcateCrab()
}
