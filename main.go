package main

import (
	"fmt"
	"s/detector"
)

func main() {

	var name string

	fmt.Print("Enter your input: ")
	fmt.Scan(&name)

	v := detector.Detector{Input: name}

	// pattern := `^[0-9]+\.[0-9]+$`
	// ok, err := regexp.MatchString(pattern, name)

	// if err != nil {
	// 	fmt.Println("Erro", err.Error())
	// }

	// fmt.Println(ok)
}
