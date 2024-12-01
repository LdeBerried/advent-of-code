package main

import (
	"fmt"
	"github.com/LdeBerried/advent-of-code/day2/hand"
	"github.com/LdeBerried/advent-of-code/internal/readfile"
	"strings"
)

func main() {
	inputContent, err := readfile.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	rawGamesArray := strings.Split(inputContent, "\n")
	ownPoints := 0

	//First part
	for _, game := range rawGamesArray {
		gameBytes := []byte(game)

		otherHand, err := hand.NewHand(gameBytes[0])
		if err != nil {
			panic(err)
		}
		ownHand, err := hand.NewHand(gameBytes[2])
		if err != nil {
			panic(err)
		}

		gameResult := ownHand.Play(otherHand)
		ownPoints = ownPoints + ownHand.GetGamePoints(gameResult)
	}

	fmt.Printf("Expected points for me: %d\n", ownPoints)

	//Second Part
	trueOwnPoints := 0
	for _, game := range rawGamesArray {
		gameBytes := []byte(game)

		otherHand, err := hand.NewHand(gameBytes[0])
		if err != nil {
			panic(err)
		}
		ownHand, err := hand.NewHandByGameResult(otherHand, gameBytes[2])
		if err != nil {
			panic(err)
		}

		gameResult := ownHand.Play(otherHand)
		trueOwnPoints = trueOwnPoints + ownHand.GetGamePoints(gameResult)
	}

	fmt.Printf("Expected true points for me: %d\n", trueOwnPoints)

}
