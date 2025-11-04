package main

import (
	"fmt"
	"regexp"
)

func main() {

	Patterns := map[string]string{
		"^[a-zA-Z_][a-zA-Z_0-9]*$":               "varables",
		"^[0-9]+$":                               "z - number",
		`^[0-9]+\.[0-9]+$`:                       "q - number",
		`^".*"$`:                                 "strings",
		"^[><=!][=]{0,1}$":                       "Operators",
		"(^if$|^else$|^for$)":                    "Commands",
		"(^int$|^float$|^string$|^bool$|^char$)": "Reserved",
		"^[(){}]{1}$":                            "Bracketes",
		"(^return$|^break$)":                     "Reserve2",
		"^[;,]{1}$":                              "Comma",
	}

	for {
		var input string

		fmt.Print("Enter your input and i guess: ")
		fmt.Scan(&input)
		y := false
		for pat, msg := range Patterns {
			ok, err := regexp.MatchString(pat, input)
			if err != nil {
				fmt.Println("🚫 regexp Error")
			}

			if ok {
				fmt.Println("✅ ", msg)
				y = true
				break

			}
		}

		if !y {
			fmt.Printf("‼️ Oops what is this? %s\n", input)
		}
	}

}
