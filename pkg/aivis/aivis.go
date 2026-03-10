package aivis

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type Aivis struct {
	c *resty.Client
	e *Engine
}

func New(baseurl string) *Aivis {
	return &Aivis{
		c: resty.New().SetBaseURL(baseurl).SetTimeout(time.Second * 10),
	}
}
