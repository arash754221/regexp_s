package fileReader

import (
	"fmt"
	"regexp"
)

func Tokenizer(btIn []byte) []string {
	re := regexp.MustCompile(`\d+\.\d+|\d+|==|!=|>=|<=|[(){};,]|[><=!+\-*/]|"[^"]*"|\w+|\S`)
	tokens := re.FindAllString(string(btIn), -1)

	for _, q := range tokens {

		fmt.Printf("%v \n", q)
	}
	return tokens
}
