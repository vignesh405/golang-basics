package main

import "fmt"

//two main functions not allowed inside same package
func main() {
	fmt.Println("hello")
}

//two functions not allowed inside same package with same name
func sample() {
	fmt.Println("hey im the duplicate")
}
