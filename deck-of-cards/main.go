package main

import (
	"./deck"
	"fmt"
)

func main() {
	cards := deck.New(deck.DefaultSort)

	// cards = deck.Shuffle(cards)

	// cards = deck.DefaultSort(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}
