package game

import (
	"sync"

	"github.com/XJIeI5/card_game/pkg/gamelogic/player"
)

type turnManager struct {
	mu            sync.Mutex
	passedPlayers map[*player.Player]bool
}

func newTurnManager() *turnManager {
	return &turnManager{
		passedPlayers: make(map[*player.Player]bool),
		mu:            sync.Mutex{},
	}
}

func (tm *turnManager) isPassed(pl *player.Player) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	return tm.passedPlayers[pl]
}

func (tm *turnManager) passFor(pl *player.Player) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.passedPlayers[pl] = true
}

func (tm *turnManager) reset() {
	clear(tm.passedPlayers)
}
