package blackjack

import (
	"math/rand"
	"time"
)

type Player struct {
	Cards []int
	Hand  int
	Score int
}

type Dealer struct {
	Cards []int
	Hand  int
	Score int
}

func (p *Player) DrawCard() {
	if !p.CheckHand() {
		return
	}

	rand.Seed(time.Now().UnixNano())

	p.Cards = append(p.Cards, p.EvaluateCard(rand.Intn(12-1)+1))
	p.Hand = GetHand(p.Cards)
}

func (p Player) CheckHand() bool {
	if p.Hand >= 21 {
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

func (d *Dealer) DrawCard() {
	if !d.CheckHand() {
		return
	}

	rand.Seed(time.Now().UnixNano())

	d.Cards = append(d.Cards, d.EvaluateCard(rand.Intn(12-1)+1))
	d.Hand = GetHand(d.Cards)
}

func (d Dealer) CheckHand() bool {
	if d.Hand >= 17 {
		return false
	}

	return true
}

func (d Dealer) EvaluateCard(card int) int {
	if d.Hand >= 11 && card == 11 || FindAce(d.Cards) && card == 11 {
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
		p.Hand, d.Hand = 0, 0
	}()

	if p.Hand > d.Hand && p.Hand <= 21 || d.Hand > 21 && p.Hand <= 21 {
		p.Score++
	} else if p.Hand == d.Hand || p.Hand > 21 && d.Hand > 21 {
		return
	} else {
		d.Score++
	}
}
