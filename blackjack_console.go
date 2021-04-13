package main

import (
	"blackjack/src"
	"fmt"
	"time"
)

var player blackjack.Player
var dealer blackjack.Dealer

func ShowMenu() {
	fmt.Println("\n\nOptions")
	fmt.Println("-----------------")
	fmt.Println("1. Hit \n2. Stand \n3. Quit")
	fmt.Print("> ")
}

func DrawCards() {
	fmt.Println("\nStarting a new round...")

	for i := 0; i < 2; i++ {
		player.DrawCard()
		dealer.DrawCard()
	}

	time.Sleep(time.Second)
}

func DealerTurn() {
	defer DrawCards()

	for dealer.CanDraw() {
		dealer.DrawCard()
		fmt.Printf("Dealer draws a %v\n", dealer.Cards[len(dealer.Cards) - 1])
		time.Sleep(time.Second)
	}

	fmt.Printf("Dealer's hand: %v\n", blackjack.GetHand(dealer.Cards))
	blackjack.TallyScore(&player, &dealer)
}

func Play() {
	DrawCards()

	var c int
	for c != 3 {
		fmt.Println("\n\nScores")
		fmt.Println("----------------")
		fmt.Printf("Player: %v\nDealer: %v", player.Score, dealer.Score)
		fmt.Printf("\nPlayer Hand: %v", blackjack.GetHand(player.Cards))
		fmt.Printf("\nDealer is showing a %v", dealer.Cards[0])

		ShowMenu()
		fmt.Scanf("%v", &c)
		fmt.Println()

		switch {
		case c == 1:
			player.DrawCard()

			if blackjack.GetHand(player.Cards) > 21 {
				fmt.Print("BUST!\n\n")
				DealerTurn()
			}

		case c == 2:
			time.Sleep(time.Second)
			DealerTurn()

		case c == 3:
			fmt.Println("\nThanks for playing!")
		}
	}
}

func main() {

	fmt.Println("-----------------------------")
	fmt.Println("|         Blackjack          |")
	fmt.Println("-----------------------------")

	Play()

	fmt.Println("\nFinal Scores")
	fmt.Println("-----------------")
	fmt.Printf("Player: %v \nDealer: %v", player.Score, dealer.Score)
}
