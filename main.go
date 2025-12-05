package main

import (
	"fmt"
	"os"
	fileReader "s/filereader"
	"s/filewriter"
	"s/validator"
)

func main() {

	mtn, err := fileReader.Read()

	if err != nil {
		fmt.Println(err)
		return
	}

	tokens := fileReader.Tokenizer(mtn)

	dataw := validator.Check(tokens)

	s := filewriter.Making(dataw)

	os.WriteFile("o.md", s, 777)

	// fmt.Printf("%+v\n", final)

	// printFile := []byte("|tokens|status|type|") +

	// printFile = append(printFile, byte("|tokens|status|type|"))

}
