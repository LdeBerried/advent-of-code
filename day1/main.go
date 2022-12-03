package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputContent, err := ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	rawCaloriesArray := strings.Split(inputContent, "\n")
	groupedCaloriesArray := make([]int, 0, 1)

	elfCalories := 0

	for _, s := range rawCaloriesArray {
		if s != "" {
			intValue, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			elfCalories += intValue
		} else {
			groupedCaloriesArray = append(groupedCaloriesArray, elfCalories)
			elfCalories = 0
		}
	}

	sort.Ints(groupedCaloriesArray)

	fmt.Printf("The elf carrying most calories is carrying %d! \n", groupedCaloriesArray[len(groupedCaloriesArray)-1])
}

func ReadFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)

	return string(fileContent), err
}
