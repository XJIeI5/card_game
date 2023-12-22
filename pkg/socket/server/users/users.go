package users

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/XJIeI5/card_game/pkg/gamelogic/card"
	"github.com/XJIeI5/card_game/pkg/gamelogic/game"
	"github.com/XJIeI5/card_game/pkg/gamelogic/player"
)

var players_list []*player.Player = make([]*player.Player, 0)

type Session struct {
	mu          sync.Mutex
	userCounter int
	users       map[int]*User
	GameSession *game.Game
}

func NewSession() *Session {
	return &Session{
		users: make(map[int]*User),
	}
}

func (s *Session) CreateUser(conn net.Conn, handler ErrorHandler) {
	if s.GameSession != nil {
		conn.Write([]byte("error: game already starts"))
		return
	}

	newPlayer := player.New(
		fmt.Sprintf("Guest%d", s.userCounter),
		card.GetStandartDeck(),
	)
	user := &User{
		id:         s.userCounter,
		connection: conn,
		session:    s,
		player:     newPlayer,
		json_mutex: GetStandartJsonMux(handler),
	}
	players_list = append(players_list, newPlayer)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.users[s.userCounter] = user
	log.Printf("Created user %d\n", s.userCounter)
	s.userCounter++
	go user.listen()
}

func (s *Session) StartGame() {
	if s.GameSession != nil {
		for _, u := range s.users {
			u.RespondeError("game already starts")
		}
		return
	}
	s.GameSession = game.New(players_list)
}

func (s *Session) Quit(index int) error {
	if _, ok := s.users[index]; !ok {
		return fmt.Errorf("error: no such user in this session")
	}
	s.users[index].connection.Close()
	delete(s.users, index)
	s.userCounter--
	log.Printf("User %d quit", index)
	return nil
}
