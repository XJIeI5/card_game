package player

import (
	"github.com/XJIeI5/card_game/pkg/gamelogic/card"
	"github.com/nuttech/bell/v2"
)

const (
	EventPlayCard   = "playcard"
	EventPlayerPass = "playerpass"
)

var SharedEvents = BellState{bell.New()}

type BellState struct {
	*bell.Events
}

func (b *BellState) Reset() {
	b.Events = bell.New()
}

type PlayCardConfig struct {
	CardIndex      int
	AsCreature     bool
	Property       *card.Property
	PeekedCreature *int
}

type PlayCardResult struct {
	cfg    PlayCardConfig
	player *Player
}

func (c *PlayCardResult) Player() *Player {
	return c.player
}

func (c *PlayCardResult) Config() PlayCardConfig {
	return c.cfg
}

type PassConfig struct {
	player *Player
}

func (c *PassConfig) Player() *Player {
	return c.player
}
