package server

import "net"

type Users map[int]*User

var (
	counter int
	users   = make(Users, 0)
)

func CreateUser(conn net.Conn) {
	user := &User{
		id:         counter,
		connection: conn,
	}
	users[counter] = user
	counter++
	go user.listen()
}
