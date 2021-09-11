package main

import "fmt"

type User struct {
	Name  string
	Email string
	age   int
}

func main() {
	vignesh := User{"Vignesh", "rvignesh12jun@gmail.com", 24}
	fmt.Printf("Simple %v\n", vignesh)
	fmt.Printf("Detail %+v\n", vignesh)
	fmt.Printf("Specific prop %v\n", vignesh.Name)

	var sam = new(User)
	sam.Name = "sam"
	sam.Email = "sam@gmail.com"
	sam.age = 22

	fmt.Printf("new method reference obj %+v\n", sam)

	var tobby = &User{"tobby", "tobby@gmail.com", 2}
	fmt.Printf("pointer %+v\n", tobby)
}
