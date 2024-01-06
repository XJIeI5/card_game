package card

type property struct {
	propertyType       PropertyType
	canBeMultiple      bool
	title              string
	description        string
	reqFoodAmount      int
	needLinkedCreature bool
}

func (p property) GetPropertyType() PropertyType {
	return p.propertyType
}

func (p property) Title() string {
	return p.title
}

func (p property) Description() string {
	return p.description
}

func (p property) ReqFoodAmount() int {
	return p.reqFoodAmount
}

func (p property) CanBeMultiple() bool {
	return p.canBeMultiple
}

// Implements

type sharpVision struct {
	property
}

func (sh sharpVision) AssignToPair(pair *Creature) {
}

func (sh sharpVision) LinkedCreature() *Creature {
	return nil
}

func newSharpVision() sharpVision {
	return sharpVision{property: property{
		propertyType:       Status,
		canBeMultiple:      false,
		title:              "Sharp Vision",
		description:        "vision is sharper",
		reqFoodAmount:      0,
		needLinkedCreature: false,
	}}
}

type fatTissue struct {
	property
}

func (ft fatTissue) AssignToPair(pair *Creature) {
}

func (ft fatTissue) LinkedCreature() *Creature {
	return nil
}

func newFatTissue() fatTissue {
	return fatTissue{property: property{
		propertyType:       Active,
		canBeMultiple:      true,
		title:              "Fat Tissue",
		description:        "many food",
		reqFoodAmount:      0,
		needLinkedCreature: false,
	}}
}
