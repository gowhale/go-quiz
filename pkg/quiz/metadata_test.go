package quiz

import (
	"fmt"
	os "os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const ()

type metadataTest struct {
	suite.Suite

	mockLoader *mockLoader
}

func (m *metadataTest) SetupTest() {
	m.mockLoader = new(mockLoader)
}

func TestMetadataTest(t *testing.T) {
	suite.Run(t, new(metadataTest))
}

func (m *metadataTest) TestGetQuestionsPass() {
	m.mockLoader.On("Open", "metadata/1/questions.json").Return(&os.File{}, nil)
	m.mockLoader.On("Close", &os.File{}).Return(nil)
	m.mockLoader.On("ReadAll", &os.File{}).Return([]byte{}, nil)
	m.mockLoader.On("Unmarshal", []byte{}).Return([]question{}, nil)
	questions, err := getQuestions(m.mockLoader)
	m.Equal([]question{}, questions)
	m.Nil(err)
}

func (m *metadataTest) TestGetQuestionsOpenError() {
	m.mockLoader.On("Open", "metadata/1/questions.json").Return(&os.File{}, fmt.Errorf("open err"))
	questions, err := getQuestions(m.mockLoader)
	m.Equal([]question{}, questions)
	m.EqualError(err, "open err")
}

func (m *metadataTest) TestGetQuestionsCloseErr() {
	m.mockLoader.On("Open", "metadata/1/questions.json").Return(&os.File{}, nil)
	m.mockLoader.On("Close", &os.File{}).Return(fmt.Errorf("close err"))
	m.mockLoader.On("ReadAll", &os.File{}).Return([]byte{}, nil)
	m.mockLoader.On("Unmarshal", []byte{}).Return([]question{}, nil)
	questions, err := getQuestions(m.mockLoader)
	m.Equal([]question{}, questions)
	m.EqualError(err, "close err")
}

func (m *metadataTest) TestGetQuestionsReadAllError() {
	m.mockLoader.On("Open", "metadata/1/questions.json").Return(&os.File{}, nil)
	m.mockLoader.On("Close", &os.File{}).Return(nil)
	m.mockLoader.On("ReadAll", &os.File{}).Return([]byte{}, fmt.Errorf("read all err"))
	questions, err := getQuestions(m.mockLoader)
	m.Equal([]question(nil), questions)
	m.EqualError(err, "read all err")
}

func (m *metadataTest) TestGetQuestionsUnmarshallError() {
	m.mockLoader.On("Open", "metadata/1/questions.json").Return(&os.File{}, nil)
	m.mockLoader.On("Close", &os.File{}).Return(nil)
	m.mockLoader.On("ReadAll", &os.File{}).Return([]byte{}, nil)
	m.mockLoader.On("Unmarshal", []byte{}).Return([]question{}, fmt.Errorf("unmarshal err"))
	questions, err := getQuestions(m.mockLoader)
	m.Equal([]question{}, questions)
	m.EqualError(err, "unmarshal err")
}

func (m *metadataTest) TestGetQuestionsUnmarshallCloseError() {
	m.mockLoader.On("Open", "metadata/1/questions.json").Return(&os.File{}, nil)
	m.mockLoader.On("Close", &os.File{}).Return(fmt.Errorf("close err"))
	m.mockLoader.On("ReadAll", &os.File{}).Return([]byte{}, nil)
	m.mockLoader.On("Unmarshal", []byte{}).Return([]question{}, fmt.Errorf("unmarshal err"))
	questions, err := getQuestions(m.mockLoader)
	m.Equal([]question{}, questions)
	m.EqualError(err, "unmarshal err close err")
}
