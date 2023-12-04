package game

import "github.com/XJIeI5/card_game/pkg/gamelogic/player"

type GamePhase int

const (
	_ GamePhase = iota
	Development
	FoodSet
	Feeding
	Extinction
)

type Game struct {
	players            []*player.Player
	currentPhase       GamePhase
	currentTurn        int
	currentFirstPlayer int
}

func New(players []*player.Player) *Game {
	game := &Game{
		players:            players,
		currentPhase:       Development,
		currentTurn:        0,
		currentFirstPlayer: 0,
	}
	return game
}

func (g *Game) NextTurn() {
	g.currentTurn++
	g.currentTurn = g.currentTurn % len(g.players)
}

func (g *Game) Run() {
	switch g.currentPhase {
	case Development:
		g.develop()
	}
}

func (g *Game) develop() {
	var savedPlayerAmount int

	for savedPlayerAmount < len(g.players) {
		currPlayer := g.players[g.currentTurn]
		if currPlayer.IsSaved() {
			g.NextTurn()
			continue
		}
		value := <-currPlayer.CardInput.IsSet
		if value {
			currPlayer.PlayCard()
		}
	}
}
