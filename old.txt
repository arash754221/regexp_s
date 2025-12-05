package main

import (
	"fmt"
	"strings"
)

func main() {

	// Patterns := map[string]string{
	// 	"^[a-zA-Z_][a-zA-Z_0-9]*$":               "variable",
	// 	"^[0-9]+$":                               "integer",
	// 	`^[0-9]+\.[0-9]+$`:                       "float",
	// 	`^".*"$`:                                 "string",
	// 	"^[><=!][=]{0,1}$":                       "operator",
	// 	"(^if$|^else$|^for$)":                    "keyword",
	// 	"(^int$|^float$|^string$|^bool$|^char$)": "type",
	// 	"^[(){}]{1}$":                            "bracket",
	// 	"(^return$|^break$)":                     "control",
	// 	"^[;,]{1}$":                              "separator",
	// }

	var query string

	fmt.Print("Enter Programm: ")
	fmt.Scan(&query)

	str := strings.Split(query, ``)
	for i, o := range str {

		fmt.Println(i, o)
	}
	// y := false
	// for pat, msg := range Patterns {
	// 	ok, err := regexp.MatchString(pat, input)
	// 	if err != nil {
	// 		fmt.Println("🚫 regexp Error")
	// 	} else {
	// 		if ok {
	// 			fmt.Println("✅ ", msg)
	// 			y = true
	// 			break

	// 		}
	// 	}
	// }

	// if !y {
	// 	fmt.Printf("‼️ Oops what is this? %s\n", input)
	// }

}
