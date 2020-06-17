package main

import (
	"fmt"
	deck "go-excercises/playing-cards/deck-package"

	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ",")
}

func (h Hand) DealerString() string {
	str := h[0].String()
	return fmt.Sprintf("%s, **HIDDEN CARD**", str)
}

func (h Hand) Score() int {
	score := 0
	for _, c := range h {
		if c.Rank == deck.Ace {
			if h.minScore() > 11 {
				score = h.minScore()
				return score

			}
			score = h.minScore() + 10
			return score
		}
	}
	score = h.minScore()
	return score
}

func (h Hand) minScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)

	var player, dealer Hand
	var c deck.Card
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			c, cards = draw(cards)
			*hand = append(*hand, c)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Print("What will you do? (h)it or (s)tand?:")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		switch input {
		case "h":
			c, cards = cards[0], cards[1:]
			player = append(player, c)
		}
	}

	pScore := player.Score()

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.minScore() != 17) {
		c, cards = draw(cards)
		dealer = append(dealer, c)
	}
	dScore := dealer.Score()

	fmt.Println("==Final Hands==")
	fmt.Println("Player:", player, "\n Score:", pScore)
	fmt.Println("Dealer:", dealer.DealerString(), "\n Score:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win")
	case dScore > pScore:
		fmt.Println("You lose")
	case pScore == dScore:
		fmt.Println("Draw")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	c, cards := cards[0], cards[1:]
	return c, cards
}
