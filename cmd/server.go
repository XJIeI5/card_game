package main

import (
	"encoding/json"
	"net"

	"github.com/XJIeI5/card_game/cmd/helper"
	"github.com/XJIeI5/card_game/pkg/socket/server/json_mux"
	"github.com/XJIeI5/card_game/pkg/socket/server/users"
	tcp_server "github.com/sylpheeed/go-tcp-socket-chat"
)

func listen(u *users.User) {
	data := json_mux.TemplateRequest{}
	buf := make([]byte, 1024)
	mux := helper.GetStandartJsonMux(func(err error) { u.RespondeError(err.Error()) })
	for {
		n, err := u.ReadConnection(buf)
		if err != nil {
			u.RespondeError("error: can't read from connection")
		}
		u.RespondeMessage(string(buf[:n]))
		if err := json.Unmarshal(buf[:n], &data); err != nil {
			u.RespondeError("undefiend responce")
			u.Quit()
		}

		mux.ServeJson(u, data)
	}
}

func main() {
	serv := tcp_server.New("localhost:9999")
	users.SetListenerForUsers(listen)
	serv.OnNewClientCallback(func(conn net.Conn) {
		users.CreateUser(conn)
	})
	serv.Listen()
}
