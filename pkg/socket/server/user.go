package server

import (
	"encoding/json"
	"fmt"
	"net"
)

type User struct {
	id         int
	name       string
	connection net.Conn
}

type RegistNewUser struct {
	Name string `json:"name"`
}

func (u *User) listen() {
	data := make(map[string]string)
	buf := make([]byte, 1024)
	for {
		n, err := u.connection.Read(buf)
		if err != nil {
			u.Responde("error: can't read from connection")
		}
		u.Responde(string(buf[:n]))
		if err := json.Unmarshal(buf[:n], &data); err != nil {
			u.Responde("error: undefiend responce")
			u.Quit()
		}

		u.name = data["name"]
		u.Responde("hello")
	}
}

func (u *User) Quit() {
	u.connection.Close()
	delete(users, u.id)
}

func (u *User) Responde(msg string) {
	fmt.Fprintf(u.connection, "%s: %s\n", u.name, msg)
}
