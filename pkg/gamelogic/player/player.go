package player

import (
	"github.com/XJIeI5/card_game/pkg/gamelogic/card"
)

type Player struct {
	Name      string
	Hand      card.Deck
	creatures []*card.Creature
	SaveInput bool
	CardInput *PlayCardConfig
	err       error
}

func New(name string, cards card.Deck) *Player {
	player := &Player{
		Name:      name,
		Hand:      cards,
		SaveInput: false,
	}
	player.creatures = make([]*card.Creature, 0)
	return player
}

func (p *Player) Error() error {
	return p.err
}

func (p *Player) AddError(err error) {
	p.err = err
}
