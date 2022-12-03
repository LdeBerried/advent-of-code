package main

import (
	"fmt"
	"os"
)

func main() {
	inputContent, err := ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(inputContent)
}

func ReadFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)

	return string(fileContent), err
}
