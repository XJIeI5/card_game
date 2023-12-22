package users

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/XJIeI5/card_game/pkg/gamelogic/player"
)

type UserListener func(*User)

type User struct {
	id         int
	name       string
	session    *Session
	player     *player.Player
	connection net.Conn
	json_mutex *JsonMux
}

func (u *User) listen() {
	data := TemplateRequest{}
	buf := make([]byte, 1024)
	mux := u.json_mutex //.GetStandartJsonMux(func(err error) { u.RespondeError(err.Error()) })
	for {
		n, err := u.ReadConnection(buf)
		if err != nil {
			u.RespondeError("error: can't read from connection")
		}
		u.RespondeMessage(string(buf[:n]))
		if err := json.Unmarshal(buf[:n], &data); err != nil {
			u.RespondeError("undefiend responce")
			u.session.Quit(u.id)
			break
		}

		mux.ServeJson(u, data)
	}
}

func (u *User) SetName(name string) {
	u.name = name
	u.player.Name = name
}

func (u *User) RespondeMessage(msg string) {
	body := make(map[string]string)
	body["name"] = u.name
	body["msg"] = msg

	responce := make(map[string]interface{})
	responce["type"] = "message"
	responce["body"] = body
	byteArr, _ := json.Marshal(responce)
	u.RespondeByteArray(byteArr)
}

func (u *User) RespondeError(err string) {
	responce := make(map[string]string)
	responce["type"] = "error"
	responce["msg"] = err
	if data, err := json.Marshal(responce); err == nil {
		fmt.Fprint(u.connection, string(data))
		return
	}
	panic("the json is incorrectly compiled")
}

func (u *User) RespondeByteArray(msg []byte) {
	fmt.Fprint(u.connection, string(msg))
}

func (u *User) ReadConnection(buf []byte) (int, error) {
	return u.connection.Read(buf)
}
