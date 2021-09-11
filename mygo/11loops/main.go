package main

import "fmt"

func main() {
	var start int = 1
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", i+start)
	}

	var arr = []string{"hey", "hello", "how are you"}
	for i := range arr {
		fmt.Printf("\n%v", arr[i])
	}

	//while like loop
	for start < 10 {
		if start > 5 {
			break
		} else if start%2 == 0 {
			start++
			continue
		}

		fmt.Printf("\n%v", start)
		if start == 5 {
			goto viglabel
		}
		// if start == 500 {
		// 	goto viglabel
		// }
		start++
	}
	fmt.Println("This will not appear in the output")
viglabel:
	{
		fmt.Println("This is viglabel")
		fmt.Println("This is the second line in the label")
	}

	fmt.Println("This is outside the label")

}
