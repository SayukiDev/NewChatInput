package aivis

type Speaker struct {
	Name        string   `json:"name"`
	SpeakerUUID string   `json:"speaker_uuid"`
	Styles      []Styles `json:"styles"`
	Version     string   `json:"version"`
}
type Styles struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (a *Aivis) Speakers() ([]Speaker, error) {
	r, err := a.c.R().Get("/speakers")
	if err != nil {
		return nil, err
	}
	var rT []Speaker
	err = parseResponse(r, &rT)
	if err != nil {
		return nil, err
	}
	return rT, nil
}
