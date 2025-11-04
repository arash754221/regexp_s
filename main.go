package main

import (
	"fmt"
	"s/detector"
)

func main() {

	var name string

	fmt.Print("Enter your input and i guess: ")
	fmt.Scan(&name)

	v := detector.Detector{Input: name}

	comma := v.Comma()
	if comma {
		fmt.Println("Comma(; or ,) detected")
	}

	command := v.Commands()
	if command {
		fmt.Println("Command (if else for) detected")
	}

	reserved2 := v.Reserve2()
	if reserved2 {
		fmt.Println("Reserved (return or break) detected")
	}

	reserved := v.Reserved()
	if reserved {
		fmt.Println("Reserved (int fload string bool char) detected")
	}

	operators := v.Operators()
	if operators {
		fmt.Println("Operators (> < = <= >= == !=) detected")
	}

	bracket := v.Bracketes()
	if bracket {
		fmt.Println("Bracket ( } ) { ( ) detected")
	}

	Qnumber := v.QNumber()
	if Qnumber {
		fmt.Println("decimal number detected")
	}

	Znumber := v.ZNumber()
	if Znumber {
		fmt.Println("integer number detected")
	}

	variable := v.Variable()
	if variable {
		fmt.Println("Variable detected")
	}

	string := v.Strings()
	if string {
		fmt.Println("String detected")
	}

}
