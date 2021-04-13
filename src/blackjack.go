package blackjack

import (
	"math/rand"
	"time"
)

type Player struct {
	Cards []int
	Score int
}

type Dealer struct {
	Cards []int
	Score int
}

func (p *Player) DrawCard() {
	if !p.CanDraw() {
		return
	}

	rand.Seed(time.Now().UnixNano())

	p.Cards = append(p.Cards, p.EvaluateCard(rand.Intn(12-1)+1))
}

func (p Player) CanDraw() bool {
	if GetHand(p.Cards) >= 21 {
		return false
	}

	return true
}

func (p Player) EvaluateCard(card int) int {
	if GetHand(p.Cards) >= 11 && card == 11 || FindAce(p.Cards) && card == 11 {
		return 1
	} else {
		return card
	}
}

func (d *Dealer) DrawCard() {
	if !d.CanDraw() {
		return
	}

	rand.Seed(time.Now().UnixNano())

	d.Cards = append(d.Cards, d.EvaluateCard(rand.Intn(12-1)+1))
}

func (d Dealer) CanDraw() bool {
	if GetHand(d.Cards) >= 17 {
		return false
	}

	return true
}

func (d Dealer) EvaluateCard(card int) int {
	if GetHand(d.Cards) >= 11 && card == 11 || FindAce(d.Cards) && card == 11 {
		return 1
	} else {
		return card
	}
}

func FindAce(cards []int) bool {
	for _, i := range cards {
		if i == 11 {
			return true
		}
	}

	return false
}

func GetHand(cards []int) int {
	sum := 0

	for _, i := range cards {
		sum += i
	}

	return sum
}

func TallyScore(p *Player, d *Dealer) {
	defer func(){
		p.Cards, d.Cards = nil, nil
	}()

	playerCards, dealerCards := GetHand(p.Cards), GetHand(d.Cards)
	if playerCards > dealerCards && playerCards <= 21 || dealerCards > 21 && playerCards <= 21 {
		p.Score++
	} else if playerCards == dealerCards || playerCards > 21 && dealerCards > 21 {
		return
	} else {
		d.Score++
	}
}
