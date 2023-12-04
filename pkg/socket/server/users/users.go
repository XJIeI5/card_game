package users

import (
	"net"

	"github.com/XJIeI5/card_game/pkg/gamelogic/card"
	"github.com/XJIeI5/card_game/pkg/gamelogic/game"
	"github.com/XJIeI5/card_game/pkg/gamelogic/player"
)

var (
	counter     int
	users       map[int]*User          = make(map[int]*User)
	players     map[int]*player.Player = make(map[int]*player.Player)
	gameSession *game.Game             = nil
	listener    UserListener           = func(u *User) { panic("listener is not set") }
)

func SetListenerForUsers(newListener UserListener) {
	listener = newListener
}

func CreateUser(conn net.Conn) {
	if gameSession != nil {
		conn.Write([]byte("error: game already starts"))
		return
	}
	user := &User{
		id:         counter,
		connection: conn,
		listener:   listener,
	}
	users[counter] = user
	counter++
	go user.listen()
}

func StartGame() {
	if gameSession != nil {
		for _, u := range users {
			u.RespondeError("game already starts")
		}
		return
	}
	players_list := make([]*player.Player, len(users))
	for i, user := range users {
		players[i] = player.New(user.name, card.GetStandartDeck())
	}
	gameSession = game.New(players_list)
	go gameSession.Run()
}

func GetPlayer(u *User) *player.Player {
	return players[u.id]
}
