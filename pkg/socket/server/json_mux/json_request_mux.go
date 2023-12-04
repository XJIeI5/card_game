package json_mux

import (
	"github.com/XJIeI5/card_game/pkg/socket/server/users"
)

type ErrorHandler func(error)
type ReqTypeHandler func(*users.User, TemplateRequest) error

type TemplateRequest struct {
	Type string                 `json:"type"`
	Body map[string]interface{} `json:"body"`
}

type JsonMux struct {
	errHandler ErrorHandler
	routs      map[string]ReqTypeHandler
}

func New(errHandler ErrorHandler) *JsonMux {
	mux := &JsonMux{errHandler: errHandler}
	mux.routs = make(map[string]ReqTypeHandler)
	return mux
}

func (jm *JsonMux) Add(req_type string, req_type_handler ReqTypeHandler) {
	jm.routs[req_type] = req_type_handler
}

func (jm *JsonMux) ServeJson(u *users.User, data TemplateRequest) {
	for key, handler := range jm.routs {
		if data.Type != key {
			continue
		}
		if err := handler(u, data); err != nil {
			jm.errHandler(err)
		}
	}
}
