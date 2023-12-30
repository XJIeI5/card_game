package main

import (
	"log"
	"net"

	"github.com/XJIeI5/card_game/pkg/socket/server/users"
	tcp_server "github.com/sylpheeed/go-tcp-socket-chat"
)

func main() {
	serv := tcp_server.New("localhost:9999")
	session := users.NewSession()
	serv.OnNewClientCallback(func(conn net.Conn) {
		handler := func(u *users.User, err error) {
			u.RespondeError(err.Error())
			log.Print(u.Name(), err)
		}
		session.CreateUser(conn, handler)
	})
	serv.Listen()
}
