package hand_test

import (
	"github.com/LdeBerried/advent-of-code/day2/hand"
	"testing"
)

func TestRetrieveHandAsRockWhenGivenA(t *testing.T) {
	h, err := hand.NewHand('A')
	if err != nil {
		t.Error(err)
	}

	if h.Gesture != "Rock" {
		t.Errorf("Expected Rock got %s", h.Gesture)
	}
}

func TestMultipleGames(t *testing.T) {
	var tests = []struct {
		ownHandCode, otherHandCode byte
		expectedResult             string
	}{
		{'A', 'A', "Tie"},
		{'A', 'X', "Tie"},
		{'B', 'X', "Win"},
		{'A', 'Z', "Win"},
		{'C', 'X', "Lose"},
	}

	for _, test := range tests {
		ownHand, err := hand.NewHand(test.ownHandCode)
		if err != nil {
			t.Error(err)
		}

		otherHand, err := hand.NewHand(test.otherHandCode)
		if err != nil {
			t.Error(err)
		}

		gameResult := ownHand.Cmp(otherHand)
		if gameResult != test.expectedResult {
			t.Errorf("Expected %s at game, got %s", test.expectedResult, gameResult)
		}
	}
}

func TestPointsByMatch(t *testing.T) {
	var tests = []struct {
		ownHandCode    byte
		gameResult     string
		expectedResult int
	}{
		{'A', "Tie", 4},
		{'X', "Tie", 4},
		{'B', "Lose", 2},
		{'Z', "Tie", 6},
		{'Y', "Win", 8},
	}

	for _, test := range tests {
		ownHand, err := hand.NewHand(test.ownHandCode)
		if err != nil {
			t.Error(err)
		}

		resultPoints := ownHand.GetGamePoints(test.gameResult)
		if resultPoints != test.expectedResult {
			t.Errorf("Expected %d points, got %d points", test.expectedResult, resultPoints)
		}

	}
}
