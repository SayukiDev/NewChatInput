package fa

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type FishAudio struct {
	*resty.Client
}

func NewFishAudio(token, model string) *FishAudio {
	return &FishAudio{
		Client: resty.New().
			SetBaseURL("https://api.fish.audio/v1/").
			SetHeaders(map[string]string{
				"Authorization": "Bearer " + token,
				"model":         model,
			}),
	}
}

func (fa *FishAudio) TTS(text string, id string) ([]byte, error) {
	r, err := fa.R().SetBody(map[string]interface{}{
		"text":         text,
		"format":       "mp3",
		"reference_id": id,
		"latency":      "balanced",
	}).SetHeader("Content-Type", "application/json").
		Post("tts")
	if err != nil {
		return nil, err
	}
	if r.StatusCode() != 200 {
		return nil, errors.New(r.Status())
	}
	return r.Body(), nil
}
