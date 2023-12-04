package card

import "fmt"

type Creature struct {
	properties []*Property
	foodAmount int
}

func NewCreature() *Creature {
	creature := &Creature{foodAmount: 1}
	creature.properties = make([]*Property, 0)
	return creature
}

func (c *Creature) ApplyProperty(prop *Property) error {
	for _, myProp := range c.properties {
		if myProp.propertyType == prop.propertyType && !prop.canBeMultiple {
			return fmt.Errorf("creature can't have two same properties")
		}
	}
	c.properties = append(c.properties, prop)
	return nil
}
