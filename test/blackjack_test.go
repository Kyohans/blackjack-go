package blackjack_test

import (
	"testing"

	blackjack "github.com/Kyohans/blackjack/src"

	"github.com/stretchr/testify/assert"
)

func setUp() (blackjack.Player, blackjack.Player) {
	return blackjack.Player{}, blackjack.Player{Dealer: true}
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

	for player.CanDraw() {
		player.DrawCard()
	}

	assert.False(t, player.CanDraw())
}

func TestEvaluatingAceWithHandOf11(t *testing.T) {
	player, _ := setUp()

	player.Cards = append(player.Cards, 11)
	player.Cards = append(player.Cards, player.EvaluateCard(11))

	assert.Equal(t, player.Cards[0] + player.Cards[1], 12)
}

func TestEvaluatingAceWithHandOf10(t *testing.T) {
	player, _ := setUp()

	player.Cards = append(player.Cards, 10)
	player.Cards = append(player.Cards, player.EvaluateCard(11))

	assert.Equal(t, player.Cards[0] + player.Cards[1], 21)
}

func TestAceIsOneIfThereIsAlreadyAnAceInHand(t *testing.T) {
	player, _ := setUp()

	player.Cards = append(player.Cards, 11)

	assert.Equal(t, player.EvaluateCard(11), 1)
}

func TestPlayerCanDrawAt17ButNotTheDealer(t *testing.T) {
	player, dealer := setUp()

	player.Cards, player.Hand = append(player.Cards, 17), 17
	dealer.Cards, dealer.Hand = append(dealer.Cards, 17), 17

	player.DrawCard()
	dealer.DrawCard()

	assert.Greater(t, player.Hand, 17)

	assert.Equal(t, dealer.Hand, 17)
}

func TestPlayerWinsWhenDealerBusts(t *testing.T) {
	player, dealer := setUp()

	dealer.Cards, dealer.Hand = append(dealer.Cards, 22), 22

	blackjack.TallyScore(&player, &dealer)

	assert.Greater(t, player.Score, dealer.Score)
}

func TestRoundEndsInATie(t *testing.T) {
	player, dealer := setUp()

	blackjack.TallyScore(&player, &dealer)

	assert.Equal(t, player.Score, dealer.Score)
}

func TestDealerWinsRound(t *testing.T) {
	player, dealer := setUp()

	for player.CanDraw() {
		player.DrawCard()
	}

	blackjack.TallyScore(&player, &dealer)

	assert.Greater(t, dealer.Score, player.Score)
}

func TestConfirmHandsAreEmptiedWhenScoresAreTallied(t *testing.T) {
	player, dealer := setUp()

	player.DrawCard()
	dealer.DrawCard()

	blackjack.TallyScore(&player, &dealer)

	assert.Equal(t, len(player.Cards), 0)
	assert.Equal(t, player.Hand, 0)

	assert.Equal(t, len(dealer.Cards), 0)
	assert.Equal(t, dealer.Hand, 0)
}
