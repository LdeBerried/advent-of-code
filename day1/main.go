package main

import (
	"fmt"
	"github.com/LdeBerried/advent-of-code/internal/readfile"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputContent, err := readfile.ReadFile("day1/input.txt")
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

	topThreeElvesCalorieCount := groupedCaloriesArray[len(groupedCaloriesArray)-3:]
	fmt.Printf("The top three elves carry a total ammount of %d calories! \n", SliceSum(topThreeElvesCalorieCount))
}

func SliceSum(sliceToSum []int) (total int) {
	for _, number := range sliceToSum {
		total += number
	}
	return
}
