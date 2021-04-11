package blackjack_test

import (
	blackjack "blackjack/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUp() (blackjack.Player, blackjack.Dealer) {
	player := blackjack.Player{0, 0}
	dealer := blackjack.Dealer{0, 0}

	return player, dealer
}

func TestCanary(t *testing.T) {
	assert.True(t, true)
}

func TestDrawCardIsBetween11And1(t *testing.T) {
	player, _ := setUp()

	player.DrawCard()

	assert.LessOrEqual(t, player.Hand, 11, 1)
}

func TestPlayerDrawsACard(t *testing.T) {
	player, _ := setUp()

	player.DrawCard()

	assert.Greater(t, player.Hand, 0)
}

func TestDrawingTwoCardsDoesntResultInBust(t *testing.T) {
	player, _ := setUp()

	player.DrawCard()
	player.DrawCard()

	assert.LessOrEqual(t, player.Hand, 21)
}

func TestNotDrawingIfHandIsABust(t *testing.T) {
	player, _ := setUp()

	player.Hand = 22

	assert.False(t, player.CheckHand())
}

func TestEvaluatingAceWithHandOf11(t *testing.T) {
	assert.Equal(t, blackjack.EvaluateCard(11, 11), 12)
}

func TestEvaluatingAceWithHandOf10(t *testing.T) {
	assert.Equal(t, blackjack.EvaluateCard(10, 11), 21)
}

func TestPlayerCanDrawAt17ButNotTheDealer(t *testing.T) {
	player, dealer := setUp()

	player.Hand, dealer.Hand = 17, 17

	player.DrawCard()
	dealer.DrawCard()

	assert.GreaterOrEqual(t, player.Hand, 17)

	assert.Equal(t, dealer.Hand, 17)
}

func TestPlayerWinsWhenDealerBusts(t *testing.T) {
	player, dealer := setUp()

	player.Hand, dealer.Hand = 14, 23

	blackjack.TallyScore(&player, &dealer)

	assert.Greater(t, player.Score, dealer.Score)
}

func TestRoundEndsInATie(t *testing.T) {
	player, dealer := setUp()

	player.Hand, dealer.Hand = 10, 10

	blackjack.TallyScore(&player, &dealer)

	assert.Equal(t, player.Score, dealer.Score)
}

func TestDealerWinsRound(t *testing.T) {
	player, dealer := setUp()

	player.Hand, dealer.Hand = 19, 21

	blackjack.TallyScore(&player, &dealer)

	assert.Greater(t, dealer.Score, player.Score)
}
