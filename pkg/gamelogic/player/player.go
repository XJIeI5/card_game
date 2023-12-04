package player

import (
	"fmt"

	"github.com/XJIeI5/card_game/pkg/gamelogic/card"
)

type Player struct {
	Name      string
	Hand      card.Deck
	creatures []*card.Creature
	*SaveInput
	*CardInput
}

func New(name string, cards []*card.Card) *Player {
	player := &Player{
		Name:      name,
		Hand:      card.NewDeck(),
		SaveInput: NewSaveInput(),
		CardInput: NewCardInput(),
	}
	player.creatures = make([]*card.Creature, 0)
	return player
}

func (p *Player) PlayCard() error {
	if value, ok := <-p.CardInput.IsSet; !(value || ok) {
		return fmt.Errorf("card input isn't set")
	}
	defer func() {
		p.CardInput = NewCardInput()
	}()

	cardIndex := p.CardInput.PlayCardConfig.CardIndex
	peekedCard := p.Hand[cardIndex]
	if p.CardInput.PlayCardConfig.AsCreature {
		p.creatures = append(p.creatures, card.NewCreature())
		return p.Hand.Remove(cardIndex)
	}

	var prop *card.Property
	if p.CardInput.PlayCardConfig.IsFirstProperty == nil {
		return fmt.Errorf("is_first_property flag isn't set")
	}

	if *p.CardInput.PlayCardConfig.IsFirstProperty {
		prop = &peekedCard.Property
	} else {
		prop = peekedCard.SecondProperty
	}

	if p.CardInput.PlayCardConfig.PeekedCreature == nil {
		return fmt.Errorf("peeked_creature field isn't set")
	}
	err := p.Hand.Remove(cardIndex)
	if err == nil {
		err = p.creatures[*p.CardInput.PlayCardConfig.PeekedCreature].ApplyProperty(prop)
	}
	return err
}
