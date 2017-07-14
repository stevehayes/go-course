package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card contains the Type (Number) and the Suit
type Card struct {
	Type string
	Suit string
}

// Deck holds an array of Card
type Deck []Card

// NewDeck will generate a new Deck
func NewDeck() (deck Deck) {
	types := []string{
		"Two", "Three", "Four", "Five", "Six", "Seven", "Eigth",
		"Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	suits := []string{"Spade", "Club", "Diamond", "Heart"}

	// Loop both types and suits to create card and add to Deck
	for t := 0; t < len(types); t++ {
		for s := 0; s < len(suits); s++ {
			card := Card{
				Type: types[t],
				Suit: suits[s],
			}
			deck = append(deck, card)
		}
	}
	return
}

// Shuffle will randomize the cards within a given deck
func Shuffle(deck Deck) Deck {
	for card := 1; card < len(deck); card++ {
		randomizer := rand.Intn(card + 1)
		if card != randomizer {
			deck[randomizer], deck[card] = deck[card], deck[randomizer]
		}
	}
	return deck
}

// Deal out a given number of cards
func Deal(deck Deck, amount int) {
	for i := 0; i < amount; i++ {
		fmt.Println(deck[i])
	}
}

// Seed using time for randomness
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	deck := NewDeck()
	Shuffle(deck)
	Deal(deck, 52)
}
