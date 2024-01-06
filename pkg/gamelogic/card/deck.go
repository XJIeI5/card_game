package card

import "fmt"

type Deck []*Card

func NewDeck() Deck {
	return make(Deck, 0)
}

func (d *Deck) Remove(index int) error {
	if index >= len(*d) || index < 0 {
		return fmt.Errorf("incorrect index")
	}
	*d = append((*d)[:index], (*d)[index+1:]...)
	return nil
}

func GetStandartDeck() Deck {
	cards := []*Card{
		{
			Property:       newSharpVision(),
			SecondProperty: nil,
		},
		{
			Property:       newFatTissue(),
			SecondProperty: nil,
		},
		{
			Property:       newSharpVision(),
			SecondProperty: newFatTissue(),
		},
	}

	return cards
}
