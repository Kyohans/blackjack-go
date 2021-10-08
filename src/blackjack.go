package blackjack

import (
	"math/rand"
	"time"
)

type Player struct {
	Cards []int
	Hand int
	Score int
	Dealer bool
}

func (p *Player) DrawCard() {
	if !p.CanDraw() {
		return
	}

	rand.Seed(time.Now().UnixNano())
	card := p.EvaluateCard(rand.Intn(12-1)+1)

	p.Cards = append(p.Cards, card)
	p.Hand += card

}

func (p Player) CanDraw() bool {
	if p.Hand >= 21 || p.Dealer && p.Hand >= 17 {
		return false
	}

	return true
}

func (p Player) EvaluateCard(card int) int {
	if p.Hand >= 11 && card == 11 || FindAce(p.Cards) && card == 11 {
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

func TallyScore(p *Player, d *Player) {
	defer func(){
		p.Cards, d.Cards = nil, nil
		p.Hand, d.Hand = 0, 0
	}()

	playerCards, dealerCards := p.Hand, d.Hand
	if playerCards > dealerCards && playerCards <= 21 || dealerCards > 21 && playerCards <= 21 {
		p.Score++
	} else if playerCards == dealerCards || playerCards > 21 && dealerCards > 21 {
		return
	} else {
		d.Score++
	}
}
