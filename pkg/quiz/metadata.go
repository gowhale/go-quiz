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

func (*local) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (*local) Close(jsonFile *os.File) error {
	return jsonFile.Close()
}

func (*local) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

func (*local) Unmarshal(data []byte) (v []question, err error) {
	return v, json.Unmarshal(data, &v)
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
		return nil, err
	}
	return l.Unmarshal(byteValue)
}
