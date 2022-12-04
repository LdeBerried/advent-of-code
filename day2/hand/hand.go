package hand

import "fmt"

type Hand struct {
	Gesture string
	Code    byte
}

func NewHand(handCode byte) (*Hand, error) {
	if handCode == 'X' || handCode == 'A' {
		return &Hand{"Rock", handCode}, nil
	} else if handCode == 'Y' || handCode == 'B' {
		return &Hand{"Paper", handCode}, nil
	} else if handCode == 'Z' || handCode == 'C' {
		return &Hand{"Scissors", handCode}, nil
	} else {
		return nil, fmt.Errorf("code unvalid")
	}

}

func (h *Hand) Play(other *Hand) (cmpCode string) {
	if h.Gesture == other.Gesture {
		cmpCode = "Tie"
	}

	if (h.Gesture == "Rock" && other.Gesture == "Scissors") ||
		(h.Gesture == "Paper" && other.Gesture == "Rock") ||
		(h.Gesture == "Scissors" && other.Gesture == "Paper") {
		cmpCode = "Win"
	} else if (h.Gesture == "Rock" && other.Gesture == "Paper") ||
		(h.Gesture == "Paper" && other.Gesture == "Scissors") ||
		(h.Gesture == "Scissors" && other.Gesture == "Rock") {
		cmpCode = "Lose"
	}
	return
}

func (h *Hand) GetGamePoints(gameResult string) int {
	pointsMap := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
		"Win":      6,
		"Tie":      3,
		"Lose":     0,
	}

	return pointsMap[h.Gesture] + pointsMap[gameResult]
}
