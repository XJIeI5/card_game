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
	canBeMultiple bool
	Title         string
	Description   string
	FoodAmount    int
}
