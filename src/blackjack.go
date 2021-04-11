package blackjack

import (
	"math/rand"
	"time"
)

type Player struct {
	Hand int
	Score int
}

type Dealer struct {
	Hand int
	Score int
}

func (p *Player) DrawCard() {
	if !p.CheckHand() {
		return
	}

	rand.Seed(time.Now().UnixNano())
	card := rand.Intn(12-1) + 1

	p.Hand = EvaluateCard(p.Hand, card)
}

func (p Player) CheckHand() bool {
	if p.Hand >= 21 {
		return false
	}

	return true
}

func (d *Dealer) DrawCard() {
	if !d.CheckHand() {
		return
	}

	rand.Seed(time.Now().UnixNano())
	card := rand.Intn(12-1) + 1

	d.Hand = EvaluateCard(d.Hand, card)
}

func (d Dealer) CheckHand() bool {
	if d.Hand >= 17 {
		return false
	}

	return true
}

func EvaluateCard(hand, card int) int {
	if hand >= 11 && card == 11 {
		return hand + 1
	} else {
		return hand + card
	}
}

func TallyScore(p *Player, d *Dealer) {
	if p.Hand > d.Hand && p.Hand <= 21 || d.Hand > 21 && p.Hand <= 21 {
		p.Score++
	} else if p.Hand == d.Hand {
		return
	} else {
		d.Score++
	}

}
