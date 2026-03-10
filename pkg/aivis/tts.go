package aivis

import "strconv"

type Mora struct {
	Text            string  `json:"text"`
	Consonant       string  `json:"consonant"`
	ConsonantLength float64 `json:"consonant_length"`
	Vowel           string  `json:"vowel"`
	VowelLength     float64 `json:"vowel_length"`
	Pitch           float64 `json:"pitch"`
}

type AccentPhrase struct {
	Moras           []Mora `json:"moras"`
	Accent          int    `json:"accent"`
	IsInterrogative bool   `json:"is_interrogative"`
}

type SynthesisOption struct {
	AccentPhrases      []AccentPhrase `json:"accent_phrases"`
	PauseMora          *Mora          `json:"pause_mora"`
	SpeedScale         float64        `json:"speedScale"`
	IntonationScale    float64        `json:"intonationScale"`
	TempoDynamicsScale float64        `json:"tempoDynamicsScale"`
	PitchScale         float64        `json:"pitchScale"`
	VolumeScale        float64        `json:"volumeScale"`
	PrePhonemeLength   float64        `json:"prePhonemeLength"`
	PostPhonemeLength  float64        `json:"postPhonemeLength"`
	PauseLength        *float64       `json:"pauseLength"`
	PauseLengthScale   float64        `json:"pauseLengthScale"`
	OutputSamplingRate int            `json:"outputSamplingRate"`
	OutputStereo       bool           `json:"outputStereo"`
	Kana               string         `json:"kana"`
}

func NewSynthesisOption() *SynthesisOption {
	return &SynthesisOption{
		AccentPhrases:      []AccentPhrase{},
		SpeedScale:         1.0,
		IntonationScale:    1.0,
		TempoDynamicsScale: 1.0,
		PitchScale:         0.0,
		VolumeScale:        1.0,
		PrePhonemeLength:   0.1,
		PostPhonemeLength:  0.1,
		PauseLength:        nil,
		PauseLengthScale:   1.0,
		OutputSamplingRate: 24000,
		OutputStereo:       false,
	}
}

func (a *Aivis) AudioQuery(text string, speaker int64) (*SynthesisOption, error) {
	res, err := a.c.R().SetQueryParams(map[string]string{
		"text":    text,
		"speaker": strconv.FormatInt(speaker, 10),
	}).Post("/audio_query")
	if err != nil {
		return nil, err
	}
	resT := new(SynthesisOption)
	err = parseResponse(res, resT)
	if err != nil {
		return nil, err
	}
	return resT, nil
}

func (a *Aivis) Synthesis(text string, speaker int64, options *SynthesisOption) ([]byte, error) {
	if options == nil {
		options = NewSynthesisOption()
	}
	res, err := a.c.R().
		SetQueryParams(map[string]string{
			"text":    text,
			"speaker": strconv.FormatInt(speaker, 10),
		}).
		ForceContentType("application/json").
		SetBody(options).
		Post("/synthesis")
	if err != nil {
		return nil, err
	}
	err = parseResponse(res, nil)
	if err != nil {
		return nil, err
	}
	return res.Body(), nil
}
