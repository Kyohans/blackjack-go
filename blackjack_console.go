package main

import (
	"fmt"
	"time"

	blackjack "github.com/Kyohans/blackjack/src"
)

var player = blackjack.Player{}
var dealer = blackjack.Player{Dealer: true}

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

	fmt.Printf("Dealer's second card is a %v\n", dealer.Cards[1])
	for dealer.CanDraw() {
		time.Sleep(time.Second)
		dealer.DrawCard()
		fmt.Printf("Dealer draws a %v\n", dealer.Cards[len(dealer.Cards) - 1])
	}

	fmt.Printf("Dealer's hand: %v\n", dealer.Hand)
	blackjack.TallyScore(&player, &dealer)
}

func Play() {
	DrawCards()

	var choice int
	for choice != 3 {
		fmt.Println("\n\nScores")
		fmt.Println("-----------------")
		fmt.Printf("Player: %v\nDealer: %v", player.Score, dealer.Score)
		fmt.Printf("\nPlayer Hand: %v", player.Hand)
		fmt.Printf("\nDealer is showing a %v", dealer.Cards[0])

		ShowMenu()
		fmt.Scanf("%v", &choice)
		fmt.Println()

		switch {
		case choice == 1:
			player.DrawCard()

			if player.Hand > 21 {
				fmt.Print("BUST!\n\n")
				DealerTurn()
			}

		case choice == 2:
			time.Sleep(time.Second)
			DealerTurn()

		case choice == 3:
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
