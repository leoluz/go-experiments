package httpsend_test

import (
	"errors"
	"net/http"
	"testing"

	httpsend "github.com/leoluz/go-experiments/bestpractices/interfaces/example2"
	"github.com/stretchr/testify/assert"
)

type httpSenderMock struct {
	method  string
	err     error
	doCount int
}

func (mock *httpSenderMock) Do(req *http.Request) (*http.Response, error) {
	mock.doCount++
	mock.method = req.Method
	if mock.err != nil {
		return nil, mock.err
	}
	resp := &http.Response{StatusCode: http.StatusOK}
	return resp, nil
}

func TestDoGet(t *testing.T) {
	t.Run("will send a get request", func(t *testing.T) {
		// given
		t.Parallel()
		mock := new(httpSenderMock)
		sender := &httpsend.HttpSend{Doer: mock}

		// when
		status, err := sender.DoGet("http://fake.url")

		// then
		assert.Nil(t, err)
		assert.NotNil(t, status)
		assert.Equal(t, http.MethodGet, mock.method)
		assert.Equal(t, 1, mock.doCount)
	})
	t.Run("will return error if url has invalid character", func(t *testing.T) {
		// given
		t.Parallel()
		mock := new(httpSenderMock)
		sender := &httpsend.HttpSend{Doer: mock}

		// when
		status, err := sender.DoGet("http://fake. url")

		// then
		assert.NotNil(t, err)
		assert.Equal(t, 0, status)
		assert.Equal(t, 0, mock.doCount)
	})
	t.Run("will return error if request fails", func(t *testing.T) {
		t.Skip()
		// given
		t.Parallel()
		mock := new(httpSenderMock)
		mock.err = errors.New("some error")
		sender := &httpsend.HttpSend{Doer: mock}

		// when
		status, err := sender.DoGet("http://fake.url")

		// then
		assert.NotNil(t, err)
		assert.Equal(t, 0, status)
		assert.Equal(t, 1, mock.doCount)
	})
}

func TestRealHttpSend(t *testing.T) {
	t.Run("will send real http request", func(t *testing.T) {
		t.Skip()
		// given
		t.Parallel()
		sender := httpsend.NewHttpSend()

		// when
		status, err := sender.DoGet("http://httpbin.org/status/200")

		// then
		assert.Nil(t, err)
		assert.Equal(t, 200, status)
	})
}
