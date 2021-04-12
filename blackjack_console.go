package main

import (
	blackjack "blackjack/src"
	"fmt"
	"time"
)

func ShowMenu() {
	fmt.Println("\n\nOptions")
	fmt.Println("-----------------")
	fmt.Println("1. Hit \n2. Stand \n3. Quit")
	fmt.Print("> ")
}

func DrawCards(player *blackjack.Player, dealer *blackjack.Dealer) {
	fmt.Println("Drawing cards...")
	for i := 0; i < 2; i++ {
		player.DrawCard()
		dealer.DrawCard()
	}
	time.Sleep(2 * time.Second)
}

func PlayRound(player *blackjack.Player, dealer *blackjack.Dealer) {
	DrawCards(player, dealer)

	var c int
	for c != 3 {
		fmt.Println("\n\nScores")
		fmt.Println("----------------")
		fmt.Printf("Player: %v\nDealer: %v", player.Score, dealer.Score)
		fmt.Printf("\nPlayer Hand: %v", blackjack.GetHand(player.Cards))

		ShowMenu()
		fmt.Scanf("%v", &c)
		fmt.Println()

		switch {
		case c == 1:
			player.DrawCard()
			if player.Hand > 21 {
				fmt.Println("BUST!\n")
				for dealer.CheckHand() {
					dealer.DrawCard()
					fmt.Println("Dealer draws...")
					time.Sleep(time.Second)
					fmt.Printf("Dealer hand: %v\n", blackjack.GetHand(dealer.Cards))
				}
				blackjack.TallyScore(player, dealer)
				DrawCards(player, dealer)
			}
		case c == 2:
			for dealer.CheckHand() {
				dealer.DrawCard()
				fmt.Println("Dealer draws...")
				time.Sleep(time.Second)
				fmt.Printf("Dealer hand: %v\n", blackjack.GetHand(dealer.Cards))
			}
			blackjack.TallyScore(player, dealer)
			DrawCards(player, dealer)
		case c == 3:

		}
	}
}

func main() {
	player := blackjack.Player{}
	dealer := blackjack.Dealer{}

	fmt.Println("-----------------------------")
	fmt.Println("|         Blackjack          |")
	fmt.Println("-----------------------------")

	PlayRound(&player, &dealer)

	fmt.Println("\nFinal Scores")
	fmt.Println("-----------------")
	fmt.Printf("Player: %v; Dealer: %v", player.Score, dealer.Score)
}
