package game

import "github.com/XJIeI5/card_game/pkg/gamelogic/player"

type Game struct {
	players []*player.Player
}

func New(players []*player.Player) *Game {
	return &Game{players: players}
}
