package main

import "fmt"

func main() {
	var score int = 34
	var p *int = &score

	fmt.Printf("P is having a value: %v %T \n", p, p)
	//nil pointer dereference SIGSEGV
	//fmt.Printf("P is having a value: %v %T", *p, *p)

	if p != nil {
		fmt.Printf("P is having a value: %v %T", *p, *p)
	} else {
		fmt.Println("Watch out..pointer is nil")
	}

}
