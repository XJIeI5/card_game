package game

import (
	"log"
	"sync"

	"github.com/XJIeI5/card_game/pkg/gamelogic/player"
	"github.com/nuttech/bell/v2"
)

type Game struct {
	players player.Players

	currPlayer int
	mux        sync.Mutex
}

func New(players []*player.Player) *Game {
	return &Game{players: players}
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
	if g.currPlayer >= len(g.players) {
		g.currPlayer = 0
	}
	log.Print("next, current ", g.currPlayer)
}

func (g *Game) Run() {
	g.Step()
	g.Run()
}

func (g *Game) Step() {
	g.develop()
	log.Print("end of turn")
	g.players.Reset()
	player.SharedEvents.Reset()
}

func (g *Game) develop() {
	var myWg sync.WaitGroup
	myWg.Add(len(g.players))

	player.SharedEvents.Listen(player.EventPlayCard, g.playCardEventHandler)

	player.SharedEvents.Listen(player.EventPlayerPass, func(message bell.Message) {
		cfg := message.(player.PassConfig)
		log.Print(cfg.Player().Name, " pass")

		myWg.Done()
	})

	player.SharedEvents.Wait()
	myWg.Wait()
}

func (g *Game) playCardEventHandler(message bell.Message) {
	pl := g.CurrentPlayer()

	if pl.IsPassed() {
		g.nextPlayer()
		return
	}

	cfg := message.(player.PlayCardResult)
	if cfg.Player() != pl {
		return
	}

	log.Print(pl.Name, " play card")
	g.nextPlayer()
}
