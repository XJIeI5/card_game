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
		Title:         "Острое зрение",
		Description:   "Далеко видит",
		FoodAmount:    0,
		canBeMultiple: false,
	}

	fatReserve := Property{
		propertyType:  FoodSave,
		Title:         "Жировой запас",
		Description:   "Может запасать еду",
		FoodAmount:    0,
		canBeMultiple: true,
	}

	cards := []*Card{
		{
			Property:       sharpEyesight,
			SecondProperty: nil,
		},
		{
			Property:       fatReserve,
			SecondProperty: nil,
		},
		{
			Property:       sharpEyesight,
			SecondProperty: &fatReserve,
		},
	}

	return cards
}
