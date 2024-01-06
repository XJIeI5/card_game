package player

import (
	"fmt"

	"github.com/XJIeI5/card_game/pkg/gamelogic/card"
	"github.com/nuttech/bell/v2"
)

type Player struct {
	Name   string
	Hand   card.Deck
	Events *bell.Events

	creatures []*card.Creature
	err       error
}

func New(name string, cards card.Deck) *Player {
	player := &Player{
		Name:   name,
		Hand:   cards,
		Events: bell.New(),
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

func (p *Player) Pass() {
	SharedEvents.Ring(EventPlayerPass, PassResult{player: p})
}

func (p *Player) PeekCard(cfg PlayCardConfig) {
	p.emitPlayCardEvent(cfg)
}

func (p *Player) PlayCard(cfg PlayCardConfig) {
	err := p.Hand.Remove(cfg.CardIndex)
	if err != nil {
		return
	}

	err = p.processCardPlay(cfg)
	if err != nil {
		return
	}

}

func (p *Player) emitPlayCardEvent(cfg PlayCardConfig) {
	res := PlayCardResult{
		cfg:    cfg,
		player: p,
	}
	SharedEvents.Ring(EventPlayCard, res)
}

func (p *Player) processCardPlay(cfg PlayCardConfig) error {
	if cfg.AsCreature {
		p.addCreature()
		return nil
	}

	if cfg.PeekedCreature == nil {
		return fmt.Errorf("no creature selected")
	}
	return p.addProperty(*cfg.PeekedCreature, cfg.Property)
}

func (p *Player) addCreature() {
	p.creatures = append(p.creatures, card.NewCreature())
}

func (p *Player) addProperty(peekedCreature int, prop card.IProperty) error {
	creature := p.creatures[peekedCreature]
	return creature.ApplyProperty(prop)
}
