package card

type Card struct {
	Property       Property
	SecondProperty *Property
}

type PropertyType int

const (
	_ PropertyType = iota
	Access
	FoodSave
)

type Property struct {
	propertyType  PropertyType
	title         string
	description   string
	foodAmount    int
	canBeMultiple bool
}
