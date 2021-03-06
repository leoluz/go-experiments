package fs_test

import (
	"errors"
	"testing"

	fs "github.com/leoluz/go-experiments/bestpractices/interfaces/example3"
	"github.com/stretchr/testify/assert"
)

type FileMock struct {
	content       string
	readFileCalls int
	err           error
}

func NewFileMock(content string, err error) *FileMock {
	return &FileMock{
		content: content,
		err:     err,
	}
}

func (m *FileMock) ReadFile(filename string) ([]byte, error) {
	m.readFileCalls++
	if m.err != nil {
		return nil, m.err
	}
	return []byte(m.content), nil
}

func TestRead(t *testing.T) {
	t.Run("read will return expected string", func(t *testing.T) {
		// given
		t.Parallel()
		expected := "some string"
		mock := NewFileMock(expected, nil)

		// when
		actual, err := fs.Read("/some/dir", mock)

		// then
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
		assert.Equal(t, 1, mock.readFileCalls)
	})
	t.Run("will return error if read fails", func(t *testing.T) {
		// given
		t.Parallel()
		mock := NewFileMock("", errors.New("some error"))

		// when
		content, err := fs.Read("/some/dir", mock)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, "", content)
		assert.Equal(t, 1, mock.readFileCalls)
	})
	t.Run("will return error if filename not informed", func(t *testing.T) {
		// given
		t.Parallel()
		mock := NewFileMock("", nil)

		// when
		content, err := fs.Read("", mock)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, "", content)
		assert.Equal(t, 0, mock.readFileCalls)
	})
}

func TestOSRead(t *testing.T) {
	t.Run("will read /etc/hosts successfully", func(t *testing.T) {
		t.Skip()
		// given
		t.Parallel()
		f := new(fs.File)

		// when
		content, err := fs.Read("/etc/hosts", f)

		// then
		assert.Nil(t, err)
		assert.NotNil(t, content)
		//assert.Equal(t, "", content)
	})
}
