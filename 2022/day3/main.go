package main

import (
	"fmt"
	"github.com/LdeBerried/advent-of-code/internal/readfile"
	"strings"
	"unicode"
)

func main() {
	inputContent, err := readfile.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	rawRucksackArray := strings.Split(inputContent, "\n")

	countPoints := 0

	for _, rucksackContent := range rawRucksackArray {
		compartment1 := strings.Split(rucksackContent[:len(rucksackContent)/2], "")
		compartment2 := strings.Split(rucksackContent[len(rucksackContent)/2:], "")

		compartment1ItemList := make(map[string]int)
		compartment2ItemList := make(map[string]int)

		for _, item := range compartment1 {
			compartment1ItemList[item] += 1
		}

		for _, item := range compartment2 {
			compartment2ItemList[item] += 1
		}

		for key, _ := range compartment1ItemList {

			if compartment2ItemList[key] > 0 {
				countPoints = countPoints + getItemPoints(key[0])
			}
		}
	}

	fmt.Println(countPoints)

}

func getItemPoints(key byte) int {
	if unicode.IsUpper(rune(key)) {
		return int(key) - 65 + 27
	} else {
		return int(key) - 96
	}
}
