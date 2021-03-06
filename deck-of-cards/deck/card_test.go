package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Diamond})
	fmt.Println(Card{Rank: Nine, Suit: Club})
	fmt.Println(Card{Rank: Jack, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Diamonds
	// Nine of Clubs
	// Jack of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 4 Suits * 13 Ranks
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	counter := 0
	for _, card := range cards {
		if card.Suit == Joker {
			counter++
		}
	}
	if counter != 3 {
		t.Error("Expected number of jokers in deck is 3. Received:", counter)
	}
}

func filterTwoAndThree(card Card) bool {
	return card.Rank == Two || card.Rank == Three
}

func TestFilters(t *testing.T) {
	cards := New(Filter(filterTwoAndThree))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("This card should not be in the deck. Received:", card)
		}
	}
}

func TestMultipleDecks(t *testing.T) {
	cards := New(MultipleDecks(2), DefaultSort)
	for i := 0; i < (len(cards) - 2); i = i + 2 {
		if cards[i] != cards[i+1] {
			t.Error("Cards are not the same:", cards[i], cards[i+1])
		}
	}
	cards = New(MultipleDecks(3))
	if len(cards) != (13 * 4 * 3) {
		t.Errorf("Expected %d cards, received %d cards", (13 * 4 * 3), len(cards))
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRank.Perm(52) should be:
	// [40 35 ...]
	shuffleOrder := []int{40, 35, 50, 0, 44, 7, 1, 16, 13, 4, 21, 12, 23, 34, 19, 11, 42, 20, 17, 48, 27, 9, 43, 46, 47, 45, 5, 49, 51, 30, 41, 26, 25, 32, 39, 28, 37, 31, 33, 10, 22, 8, 6, 29, 36, 18, 14, 2, 15, 3, 38, 24}
	shuffleRand = rand.New(rand.NewSource(0))
	cards := New(Shuffle)
	originalCards := New()
	for i, card := range cards {
		if card != originalCards[shuffleOrder[i]] {
			t.Errorf("Expected %s, received %s", originalCards[shuffleOrder[i]], card)
		}
	}
}
