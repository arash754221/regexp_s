package fileReader

import (
	"errors"
	"os"
)

func Read() ([]byte, error) {
	bt, err := os.ReadFile("./input.s")
	if err != nil {
		return nil, err
	}
	if len(bt) == 0 {
		return nil, errors.New("input file is empty")
	}
	return bt, nil

}
