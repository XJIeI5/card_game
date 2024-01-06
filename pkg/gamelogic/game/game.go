package game

import (
	"log"
	"sync"
	"time"

	"github.com/XJIeI5/card_game/pkg/gamelogic/player"
	"github.com/nuttech/bell/v2"
)

type Game struct {
	players []*player.Player
	food    *FoodSupply
	turn    turnManager

	currPlayer int
	mux        sync.Mutex
}

func New(players []*player.Player) *Game {
	return &Game{players: players, turn: *newTurnManager()}
}

func (g *Game) Run() {
	g.MakeTurn()
	g.Run()
}

func (g *Game) MakeTurn() {
	g.develop()
	g.setFoodSupply()
	// g.nutrition()
	log.Print("end of turn")
}

func (g *Game) nutrition() {
	defer g.resertEventSystem()

	player.SharedEvents.Listen(player.EventFeedFromSupply, func(message bell.Message) {
		// feedCfg := message.(player.FeedFromSupplyResult)

	})

	player.SharedEvents.Wait()
}

func (g *Game) setFoodSupply() {
	g.food = newFoodSupply(len(g.players))
}

func (g *Game) develop() {
	defer g.resertEventSystem()

	var wg sync.WaitGroup
	wg.Add(len(g.players))

	player.SharedEvents.Listen(player.EventPlayCard, func(message bell.Message) {
		g.playCardEventHandler(message)
	})

	player.SharedEvents.Listen(player.EventPlayerPass, func(message bell.Message) {
		cfg := message.(player.PassResult)
		time.Sleep(1 * time.Millisecond)
		if g.canPass(message) {
			log.Print(cfg.Player().Name, " passes")
			g.turn.passFor(cfg.Player())
			g.nextPlayer()
			wg.Done()
		}
	})

	player.SharedEvents.Wait()
	wg.Wait()
}

func (g *Game) resertEventSystem() {
	g.turn.reset()
	player.SharedEvents.Reset()
}

func (g *Game) canPass(message bell.Message) bool {
	passCfg := message.(player.PassResult)
	if g.turn.isPassed(passCfg.Player()) {
		return false
	}
	return g.CurrentPlayer() == passCfg.Player()
}

func (g *Game) playCardEventHandler(message bell.Message) {
	pl := g.CurrentPlayer()

	if g.turn.isPassed(pl) {
		g.nextPlayer()
		return
	}

	cfg := message.(player.PlayCardResult)
	if cfg.Player() != pl {
		return
	}

	log.Print(pl.Name, " plays card")
	pl.PlayCard(cfg.Config())
	g.nextPlayer()
}

func (g *Game) CurrentPlayer() *player.Player {
	g.mux.Lock()
	defer g.mux.Unlock()
	return g.players[g.currPlayer]
}

func (g *Game) nextPlayer() {
	g.mux.Lock()
	defer g.mux.Unlock()
	g.currPlayer++
	g.currPlayer %= len(g.players)
	if len(g.turn.passedPlayers) == len(g.players) {
		return
	}
	for g.turn.isPassed(g.players[g.currPlayer]) {
		g.currPlayer++
		g.currPlayer %= len(g.players)
	}
}
