package main

import "fmt"

//not allowed outside function
//jwtToken := 3000

//variable or const name that starts with caps are public
const LoginToken string = "hfgjshgf"

func main() {
	var username string = "Vignesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.12345654323456543234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var largeFloat float64 = 255.12345654323456543234
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	//default values and some aliases

	var anotherVariable int

	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type

	var website = "github.com/vignesh405"

	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style

	numberOfUsers := 3

	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	var var1, var2 string = "var1", "var2"
	fmt.Println(var1, var2)
	//not allowed
	//var var3 int, var4 string = 3,"hey"

	var (
		spiderman  = "Iam spiderman"
		age        = 18
		powers     = "all powers"
		otherstuff = "others"
	)

	fmt.Println(spiderman, age, powers, otherstuff)

}
