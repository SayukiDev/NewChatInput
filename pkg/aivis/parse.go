package aivis

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
)

type errorRes struct {
	Detail []struct {
		Loc  []string `json:"loc"`
		Msg  string   `json:"msg"`
		Type string   `json:"type"`
	} `json:"detail"`
}

func parseResponse(res *resty.Response, resT any) (err error) {
	if res.StatusCode() == 422 {
		e := new(errorRes)
		err := json.Unmarshal(res.Body(), e)
		if err != nil {
			return err
		}
		msg := ""
		for _, v := range e.Detail {
			msg += v.Msg + "\n"
		}
		return errors.New(msg)
	}
	if resT == nil {
		return
	}
	if res.StatusCode() != 200 {
		return errors.New(res.Status())
	}
	return json.Unmarshal(res.Body(), resT)
}
