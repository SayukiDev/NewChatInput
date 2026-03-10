package pages

type Input struct {
	*content
	Startup func(c *content)
}

func NewInput() *Input {
	i := &Input{}
	i.Startup = i.startup
	return i
}

func (i *Input) startup(srv *content) {
	i.content = srv
}

func (i *Input) SendMessage(message string) error {
	return i.srv.ChatBox.SendChatboxMsg(message, i.srv.Option.TTS, false)
}

func (i *Input) SetFullInputMode(mode bool) {
	if mode {
		i.content.reg.app.SetSizeRatio(5)
	} else {
		i.content.reg.app.SetSizeRatio(1.4)
	}
}

func (i *Input) SetTyping(on bool) {
	i.srv.ChatBox.SetTyping(on)
}
