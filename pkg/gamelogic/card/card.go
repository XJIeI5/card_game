package card

type Card struct {
	Property       IProperty
	SecondProperty IProperty
}

type PropertyType int

const (
	_ PropertyType = iota
	Status
	Active
	NutritionType
	Paired
)

// type Property struct {
// 	propertyType   PropertyType
// 	canBeMultiple  bool
// 	Title          string
// 	Description    string
// 	FoodAmount     int
// 	LinkedCreature *Creature

// 	AssignToPair func(pair *Creature)
// }

type IProperty interface {
	GetPropertyType() PropertyType
	Title() string
	Description() string
	ReqFoodAmount() int
	CanBeMultiple() bool
	LinkedCreature() *Creature

	AssignToPair(pair *Creature)
}
