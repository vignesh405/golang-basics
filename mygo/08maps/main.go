package main

import "fmt"

func main() {
	sampleMap := make(map[string]string)
	sampleMap["Vignesh"] = "Nallavan vallavan naalum therinjavan"
	sampleMap["Praveen"] = "KK KP LK BS"
	sampleMap["Naga Karthick"] = "Loosu Paithyam Naaga Paambu"
	sampleMap["Charavanaa"] = "Valandhu kettavan pro gamer and programmer"
	// fmt.Println(sampleMap)

	for k, v := range sampleMap {
		fmt.Println("Name is ", k)
		fmt.Println("value is", v)
	}

}
