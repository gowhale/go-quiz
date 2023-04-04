package quiz

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//go:generate go run github.com/vektra/mockery/cmd/mockery -name loader -inpkg --filename loader_mock.go
type loader interface {
	Open(name string) (*os.File, error)
	Close(jsonFile *os.File) error
	ReadAll(r io.Reader) ([]byte, error)
	Unmarshal(data []byte) ([]question, error)
}

type local struct{}

func (l *local) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (l *local) Close(jsonFile *os.File) error {
	return jsonFile.Close()
}

func (l *local) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

func (l *local) Unmarshal(data []byte) ([]question, error) {
	var v []question
	err := json.Unmarshal(data, &v)
	return v, err
}

func getQuestions(l loader) (questions []question, err error) {
	jsonFile, err := l.Open("metadata/1/questions.json")
	if err != nil {
		return []question{}, err
	}

	defer func() {
		if err2 := l.Close(jsonFile); err2 != nil {
			if err == nil {
				err = err2
			} else {
				err = fmt.Errorf("%s %s", err, err2)
			}
		}
	}()

	byteValue, err := l.ReadAll(jsonFile)
	if err != nil {
		return
	}
	return l.Unmarshal(byteValue)
}
