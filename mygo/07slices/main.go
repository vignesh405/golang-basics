package main

import "fmt"

func main() {
	things := []string{"hey", "hello", "ok", "alright"}
	fmt.Println(things)

	things = append(things, "i missed it")
	fmt.Println(things)

	things = things[1:]
	fmt.Println(things)

	heroes := make([]string, 3)
	var p *[]string = &heroes
	heroes[0] = "thor"
	heroes[1] = "ironman"
	heroes[2] = "spiderman"

	fmt.Println(heroes)
	fmt.Println("pointer to slice")
	fmt.Println(*p)

	heroes = append(heroes, "deadpool")
	fmt.Println(heroes)
	fmt.Println(cap(heroes))

	fmt.Println("pointer to slice")
	heroesCopy := make([]string, 6)
	var q *[]string = &heroesCopy
	fmt.Println(*q)
	copy(heroesCopy, heroes)
	fmt.Println("heroes", heroes)
	fmt.Println("heroesCopy", heroesCopy)
	fmt.Println("p pointer", *p)
	fmt.Println("q pointer", *q)
	fmt.Printf("p pointer ref %v \n", p)
	fmt.Printf("q pointer ref %v \n", q)
	heroesCopy = heroesCopy[1:2]
	fmt.Println(heroesCopy)
	fmt.Println(*p)
	fmt.Println(*q)

}
