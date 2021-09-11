package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the 3rd module of the series"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out Pizza")

	//comma ok syntax or error ok syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

	var myString string
	fmt.Scanln(&myString)
	fmt.Println(myString)

}
