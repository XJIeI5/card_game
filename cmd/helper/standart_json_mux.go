package helper

import (
	"fmt"

	"github.com/XJIeI5/card_game/pkg/gamelogic/player"
	"github.com/XJIeI5/card_game/pkg/socket/server/json_mux"
	"github.com/XJIeI5/card_game/pkg/socket/server/users"
)

func GetStandartJsonMux(errHandler json_mux.ErrorHandler) *json_mux.JsonMux {
	mux := json_mux.New(errHandler)
	mux.Add("regist", handleRegistNewUser)
	mux.Add("start_game", handleStartNewGame)
	mux.Add("play_card", handleCardPlay)
	return mux
}

func handleRegistNewUser(u *users.User, data json_mux.TemplateRequest) error {
	name, ok := data.Body["name"].(string)
	if !ok {
		return fmt.Errorf("body hasn't name field")
	}

	u.SetName(name)
	u.RespondeMessage("hello")
	return nil
}

func handleStartNewGame(u *users.User, data json_mux.TemplateRequest) error {
	users.StartGame()
	return nil
}

func handleCardPlay(u *users.User, data json_mux.TemplateRequest) error {
	cardIndex, ok := data.Body["card_index"].(int)
	if !ok {
		return fmt.Errorf("body hasn't card index field")
	}
	asCreature, ok := data.Body["as_creature"].(bool)
	if !ok {
		return fmt.Errorf("body hasn't as creature flag")
	}
	upperProperty, propOk := data.Body["upper_prop"].(bool)
	peeked_creature, creatureOk := data.Body["creature_index"].(int)

	config := player.PlayCardConfig{
		CardIndex:  cardIndex,
		AsCreature: asCreature,
	}
	if propOk {
		config.IsFirstProperty = &upperProperty
	}
	if creatureOk {
		config.PeekedCreature = &peeked_creature
	}
	player := users.GetPlayer(u)
	player.PeekCard(config)
	player.PlayCard()

	return nil
}
