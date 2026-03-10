package tts

import (
	malgo2 "ChatInput/pkg/malgo"
)

type Device struct {
	Name string
	Id   string
}

func (t *TTS) GetDevices() ([]Device, error) {
	ds, err := malgo2.GetDevices()
	if err != nil {
		return nil, err
	}

	newDs := make([]Device, len(ds))
	for i, d := range ds {
		newDs[i] = Device{
			Name: d.Name(),
			Id:   d.ID.String(),
		}
	}
	return newDs, nil
}
