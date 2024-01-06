package player

import (
	"github.com/XJIeI5/card_game/pkg/gamelogic/card"
	"github.com/nuttech/bell/v2"
)

const (
	EventPlayCard        = "playcard"
	EventCardCanBePlayed = "canplaycard"
	EventPlayerPass      = "playerpass"
	EventFeedFromSupply  = "feedsupply"
	EventCreatureIsFed   = "creatureisfed"
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
	Property       card.IProperty
	PeekedCreature *int
	PairedCreature *int
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

type PassResult struct {
	player *Player
}

func (p *PassResult) Player() *Player {
	return p.player
}

type FeedFromSupplyResult struct {
	player   *Player
	creature *card.Creature
}

func (f *FeedFromSupplyResult) Player() *Player {
	return f.player
}

func (f *FeedFromSupplyResult) Creature() *card.Creature {
	return f.creature
}
