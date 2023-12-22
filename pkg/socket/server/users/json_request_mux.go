package users

type ErrorHandler func(*User, error)
type ReqTypeHandler func(*User, TemplateRequest) error

type TemplateRequest struct {
	Type string                 `json:"type"`
	Body map[string]interface{} `json:"body"`
}

type JsonMux struct {
	errHandler ErrorHandler
	routs      map[string]ReqTypeHandler
}

func NewJsonMux(errHandler ErrorHandler) *JsonMux {
	mux := &JsonMux{errHandler: errHandler}
	mux.routs = make(map[string]ReqTypeHandler)
	return mux
}

func (jm *JsonMux) Add(req_type string, req_type_handler ReqTypeHandler) {
	jm.routs[req_type] = req_type_handler
}

func (jm *JsonMux) ServeJson(u *User, data TemplateRequest) {
	for key, handler := range jm.routs {
		if data.Type != key {
			continue
		}
		if err := handler(u, data); err != nil {
			jm.errHandler(u, err)
		}
	}
}
