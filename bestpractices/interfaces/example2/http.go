package httpsend

import (
	"fmt"
	"net/http"
	"time"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpSend struct {
	Doer
}

func NewHttpSend() *HttpSend {
	c := &http.Client{Timeout: time.Second * 10}
	return &HttpSend{
		Doer: c,
	}
}

func (h *HttpSend) DoGet(url string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := h.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %s", err)
	}
	return resp.StatusCode, nil
}
