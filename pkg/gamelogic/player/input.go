package player

type SaveInput struct {
	isSaved bool
}

func NewSaveInput() *SaveInput {
	return &SaveInput{}
}

func (si *SaveInput) Save() {
	si.isSaved = true
}

func (si *SaveInput) IsSaved() bool {
	return si.isSaved
}

type CardInput struct {
	IsSet chan bool
	PlayCardConfig
}

func NewCardInput() *CardInput {
	cardInput := &CardInput{}
	cardInput.IsSet = make(chan bool)
	return cardInput
}

type PlayCardConfig struct {
	CardIndex       int
	AsCreature      bool
	IsFirstProperty *bool
	PeekedCreature  *int
}

func (ci *CardInput) PeekCard(config PlayCardConfig) {
	ci.PlayCardConfig = config
	ci.IsSet <- true
}
