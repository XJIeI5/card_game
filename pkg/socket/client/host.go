package main

import "C"

import (
	"net"

	"github.com/XJIeI5/card_game/pkg/socket/server"
	tcp_server "github.com/sylpheeed/go-tcp-socket-chat"
)

//export Host
func Host(address string) {
	serv := tcp_server.New(address)
	serv.OnNewClientCallback(func(conn net.Conn) {
		server.CreateUser(conn)
	})
	serv.Listen()
}

func main() {}
