package game

import (
	"math/rand"
)

type FoodSupply struct {
	foodAmount int
}

func newFoodSupply(playerAmount int) *FoodSupply {
	foodFabric := map[int]int{
		2: throwDice() + 2,
		3: throwDice() + throwDice(),
		4: throwDice() + throwDice() + 2,
	}
	return &FoodSupply{foodAmount: foodFabric[playerAmount]}
}

func throwDice() int {
	return rand.Intn(6) + 1
}
