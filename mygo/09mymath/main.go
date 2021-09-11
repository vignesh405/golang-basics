package main

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to math in golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4.5

	fmt.Println("The sum is: ", float64(mynumberOne)+mynumberTwo)

	//random number using math/rand
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(5) + 1)

	//random number using crypto

	myRandomNum, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
