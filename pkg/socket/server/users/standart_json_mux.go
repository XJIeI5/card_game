package users

import (
	"encoding/json"
	"fmt"

	"github.com/XJIeI5/card_game/pkg/gamelogic/card"
)

func GetStandartJsonMux(errHandler ErrorHandler) *JsonMux {
	mux := NewJsonMux(errHandler)
	mux.Add("regist", handleRegistNewUser)
	mux.Add("start_game", handleStartNewGame)
	mux.Add("play_card", handleCardPlay)
	return mux
}

func handleRegistNewUser(u *User, data TemplateRequest) error {
	name, ok := data.Body["name"].(string)
	if !ok {
		return fmt.Errorf("body hasn't name field")
	}

	u.SetName(name)
	u.RespondeMessage("hello")
	return nil
}

func handleStartNewGame(u *User, data TemplateRequest) error {
	u.session.StartGame()

	player := u.player
	body := make([]card.Card, 0, len(player.Hand))
	for _, card := range player.Hand {
		body = append(body, *card)
	}

	responce := make(map[string]interface{})
	responce["type"] = "set_hand"
	responce["body"] = body
	msg, err := json.Marshal(responce)
	if err != nil {
		return err
	}
	u.RespondeByteArray(msg)

	return nil
}

func handleCardPlay(u *User, data TemplateRequest) error {
	cardIndexValue, ok := data.Body["card_index"].(float64)
	if !ok {
		return fmt.Errorf("body hasn't card index field")
	}
	fmt.Println(cardIndexValue)

	return nil
}
