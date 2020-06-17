package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Heart, Rank: Ace})

	//Output:
	//Ace of Hearts
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("wrong number of cards")
	}
}

func TestJoker(t *testing.T) {
	count := 0
	for _, card := range New(Jokers(2)) {
		if card.Suit == Joker {
			count++
		}
	}
	if count != 2 {
		t.Error("expected 2 but got", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		if card.Rank == Ace || card.Rank == Three {
			return true
		}
		return false
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == 1 || card.Rank == 3 {
			t.Error("2 or 3 should be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	numDecks := 3
	cards := New(Deck(numDecks))
	if len(cards) != 52*numDecks {
		t.Error("expected", 52*numDecks, "got", len(cards))
	}
}
