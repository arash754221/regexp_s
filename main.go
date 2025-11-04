package main

import (
	"fmt"
	"regexp"
)

func main() {

	Patterns := map[string]string{
		"^[a-zA-Z_][a-zA-Z_0-9]*$":               "variable",
		"^[0-9]+$":                               "integer",
		`^[0-9]+\.[0-9]+$`:                       "float",
		`^".*"$`:                                 "string",
		"^[><=!][=]{0,1}$":                       "operator",
		"(^if$|^else$|^for$)":                    "keyword",
		"(^int$|^float$|^string$|^bool$|^char$)": "type",
		"^[(){}]{1}$":                            "bracket",
		"(^return$|^break$)":                     "control",
		"^[;,]{1}$":                              "separator",
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
			} else {
				if ok {
					fmt.Println("✅ ", msg)
					y = true
					break

				}
			}
		}

		if !y {
			fmt.Printf("‼️ Oops what is this? %s\n", input)
		}
	}

}
