package users

import (
	"encoding/json"
	"fmt"
	"net"
)

type UserListener func(*User)

type User struct {
	id         int
	name       string
	connection net.Conn
	listener   UserListener
}

func (u *User) listen() {
	u.listener(u)
}

func (u *User) SetListener(listener UserListener) {
	u.listener = listener
}

func (u *User) Quit() {
	u.connection.Close()
	delete(users, u.id)
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) RespondeMessage(msg string) {
	fmt.Fprintf(u.connection, "%s: %s\n", u.name, msg)
}

func (u *User) RespondeError(err string) {
	responce := make(map[string]string)
	responce["type"] = "error"
	responce["msg"] = err
	if data, err := json.Marshal(responce); err == nil {
		fmt.Fprint(u.connection, data)
		return
	}
	panic("the json is incorrectly compiled")
}

func (u *User) ReadConnection(buf []byte) (int, error) {
	return u.connection.Read(buf)
}
