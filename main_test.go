package main

import (
	"fmt"
	"go-excercises/01playing-cards/deck-package"

	"testing"
)

func TestDealer(t *testing.T) {
	cards := deck.New(deck.Deck(3), deck.Shuffle)

	var dealer Hand

	dealer = []deck.Card{
		{Suit: deck.Heart,
			Rank: deck.Ace,
		},
		{Suit: deck.Club,
			Rank: deck.Six,
		},
	}

	var c deck.Card
	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.minScore() != 17) {
		c, cards = draw(cards)
		dealer = append(dealer, c)
	}

	fmt.Println("# of Cards", len(dealer))
	fmt.Println("Dealer's Cards:", dealer)

	if len(dealer) == 2 {
		t.Error("Dealer didn't draw the card")
	}
}
