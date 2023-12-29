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

func (u *User) Name() string {
	return u.name
}

func (u *User) RespondeMessage(msg string) {
	body := make(map[string]string)
	body["name"] = u.name
	body["text"] = msg

	u.RespondeStatement("message", body)
}

func (u *User) RespondeError(err string) {
	body := make(map[string]string)
	body["error"] = err

	u.RespondeStatement("error", body)
}

func (u *User) RespondeStatement(msgType string, body map[string]string) {
	responce := make(map[string]interface{})
	responce["type"] = msgType
	responce["body"] = body
	byteArr, err := json.Marshal(responce)
	if err != nil {
		panic("incorrect formed json")
	}
	u.RespondeByteArray(byteArr)
}

func (u *User) RespondeByteArray(msg []byte) {
	fmt.Fprint(u.connection, string(msg))
}

func (u *User) ReadConnection(buf []byte) (int, error) {
	return u.connection.Read(buf)
}
