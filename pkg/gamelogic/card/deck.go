package card

type Deck []*Card

func NewDeck() Deck {
	return make(Deck, 0)
}

func (d *Deck) Remove(index int) error {
	*d = append((*d)[:index], (*d)[index+1:]...)
	return nil
}

func GetStandartDeck() Deck {
	sharpEyesight := Property{
		propertyType:  Access,
		title:         "Острое зрение",
		description:   "Далеко видит",
		foodAmount:    0,
		canBeMultiple: false,
	}

	fatReserve := Property{
		propertyType:  FoodSave,
		title:         "Жировой запас",
		description:   "Может запасать еду",
		foodAmount:    0,
		canBeMultiple: true,
	}

	cards := []*Card{
		&Card{
			Property:       sharpEyesight,
			SecondProperty: nil,
		},
		&Card{
			Property:       fatReserve,
			SecondProperty: nil,
		},
		&Card{
			Property:       sharpEyesight,
			SecondProperty: &fatReserve,
		},
	}

	return cards
}
