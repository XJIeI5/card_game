package card

import "fmt"

type Creature struct {
	properties    []IProperty
	foodAmount    int
	maxFoodAmount int
}

func NewCreature() *Creature {
	creature := &Creature{foodAmount: 1}
	creature.properties = make([]IProperty, 0)
	return creature
}

func (c *Creature) ApplyProperty(prop IProperty) error {
	samePropErr := fmt.Errorf("creature can't have two same properties")

	for _, myProp := range c.properties {
		if myProp == prop && !prop.CanBeMultiple() {
			return samePropErr
		}
	}

	if prop.LinkedCreature() != nil && prop.LinkedCreature().Has(prop) {
		return samePropErr
	}

	c.properties = append(c.properties, prop)
	prop.AssignToPair(prop.LinkedCreature())
	return nil
}

func (c *Creature) Has(prop IProperty) bool {
	for _, myProp := range c.properties {
		if myProp == prop {
			return true
		}
	}
	return false
}

func (c *Creature) Feed(foodAmount int) bool {

	return false
}
