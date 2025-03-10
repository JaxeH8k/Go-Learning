package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"time"
)

type Card struct {
	Value int
	Suit  string
	Name  string
}

type Hand struct {
	Cards []Card
}

var player1 Hand
var player2 Hand

func newDeck() []Card {
	suits := []string{"♣️", "♥️", "♠️", "♦️"}
	values := []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"Jack",
		"Queen",
		"King",
		"Ace",
	}

	var cards []Card
	for _, suit := range suits {
		for i, val := range values {
			cards = append(cards, Card{
				Value: i,
				Name:  val,
				Suit:  suit,
			})
		}
	}

	return cards
}

// pretty display for cards
func (c Card) String() string {
	return fmt.Sprintf("%s : %s", c.Suit, c.Name)
}

func dealCards(deck []Card) {
	// player1 and player2 are global vars, no return needed.
	for i, card := range deck {
		if (i+1)%2 == 0 {
			player1.Cards = append(player1.Cards, card)
		} else {
			player2.Cards = append(player2.Cards, card)
		}
	}
}

func war() (string, []Card) {
	fmt.Println("War has broke out!")
	// simulate war
	isConflict := 1
	var winner string

	p1 := []Card{}
	p2 := []Card{}

	p1CDraw := 4 // amount of cards to involved to play the war round
	p2CDraw := 4 // this is dynamic in event they don't have enough cards to continue
	round := 0

	for isConflict == 1 {
		round++
		fmt.Println("Round ", round)
		// draw 4 cards from each player, 4th will be comparisson
		// if players still don't have enough for the war play, use what they have
		if len(player1.Cards) < 4 {
			fmt.Println("Player 1 card count ", len(player1.Cards))
			if len(player1.Cards) == 0 {
				// player loses
				log.Fatal("Player 2 Wins, Player 1 died in War")
			}
			// possible bug - if player has 0 cards, have to call the game
			// take the rest of their cards
			p1 = append(p1, player1.Cards[0:]...)
			player1.Cards = []Card{}
		} else {
			// draw 4 cards
			p1 = append(p1, player1.Cards[0:p1CDraw]...) // slice out of range bug
			player1.Cards = player1.Cards[p1CDraw:]
		}
		if len(player2.Cards) < 4 {
			fmt.Println("player 2 card count ", len(player2.Cards))
			if len(player2.Cards) == 0 {
				log.Fatal("Player 1 Wins, Player 2 died in War")
			}
			// same for player2 w/rt bug
			//p2CDraw = len(player2.Cards) - 1
			p2 = append(p2, player2.Cards[0:]...) // take rest of their cards
			player2.Cards = []Card{}
		} else {
			// draw 4 cards
			p2 = append(p2, player2.Cards[0:p2CDraw]...) // slice out of range bug
			player2.Cards = player2.Cards[p2CDraw:]
		}
		// draw a 4th card from each player and compare value
		if p1[len(p1)-1].Value > p2[len(p2)-1].Value {
			// declare winner, end conflict
			fmt.Printf("P1 %d | P2 %d\n", p1[len(p1)-1].Value, p2[len(p2)-1].Value)
			winner = "p1"
			isConflict = 0

		}
		if p1[len(p1)-1].Value < p2[len(p2)-1].Value {
			fmt.Printf("P1 %d | P2 %d\n", p1[len(p1)-1].Value, p2[len(p2)-1].Value)
			winner = "p2"
			isConflict = 0
		}
		// if values match we just repeate until a player wins.
	}

	fmt.Println("War won by ", winner)
	// return winner and the cards to play()
	// combine slices of the draw cards to return along with the winner
	winning := []Card{}
	winning = append(winning, p1...)
	winning = append(winning, p2...)
	return winner, winning
}

func play() {
	p1 := player1.Cards[0]
	p2 := player2.Cards[0]
	// remove the cards from their Hands
	player1.Cards = player1.Cards[1:]
	player2.Cards = player2.Cards[1:]
	var winner string
	var cards []Card

	if p1.Value > p2.Value {
		// player 1 wins, move both cards to his discard.
		player1.Cards = append(player1.Cards, p1, p2)
	}
	if p1.Value < p2.Value {
		// player 2 wins
		player2.Cards = append(player2.Cards, p1, p2)
	}
	if p1.Value == p2.Value {
		winner, cards = war()
		if winner == "p1" {
			player1.Cards = append(player1.Cards, cards...)
			player1.Cards = append(player1.Cards, p1, p2)
		}
		if winner == "p2" {
			player2.Cards = append(player2.Cards, cards...)
			player2.Cards = append(player2.Cards, p1, p2)
		}
	}
}

func main() {
	deck := newDeck()
	// // // Shuffle New Deck // // //
	// seed a random num generator
	seed := time.Now().UnixNano()
	// create a new PCG source with seed
	pcg := rand.NewPCG(uint64(seed), uint64(seed))

	// create a new rand instance using the pcg source
	r := rand.New(pcg)

	r.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	// Deck has been shuffled

	dealCards(deck)

	//fmt.Println(player1.Cards)
	//fmt.Println(player2.Cards)

	winner := ""
	for winner == "" {
		play()
		// determine winner if either of the player is out of cards
		if len(player1.Cards) == 0 {
			winner = "Player 2"
		}
		if len(player2.Cards) == 0 {
			winner = "Player 1"
		}
	}

	fmt.Println(winner, " has won the game!")
	fmt.Println("Player 1 Cards: ", len(player1.Cards))
	fmt.Println("Player 2 Cards: ", len(player2.Cards))
}
