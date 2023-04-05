package quiz

import (
	"fmt"
	"testing"

	fruit "github.com/gowhale/go-test-data/pkg/fruits"
	"github.com/stretchr/testify/suite"
)

const ()

type quesstionTest struct {
	suite.Suite

	exampleQuestion question
	mockInput       *mockInputter
}

func (m *quesstionTest) SetupTest() {
	m.exampleQuestion = question{
		Question: fruit.Apple,
		Options: []string{
			fruit.Apricot,
			fruit.Avocado,
			fruit.Blueberry,
			fruit.Lemon,
		},
		Answer: fruit.Lemon,
	}
	m.mockInput = new(mockInputter)
}

func TestQuestionTest(t *testing.T) {
	suite.Run(t, new(quesstionTest))
}

func (m *quesstionTest) TestGetInputPass() {
	m.mockInput.On("getInput").Return("1", nil)
	i, err := getInput(m.exampleQuestion, m.mockInput)
	m.Nil(err)
	m.Equal(0, i)
}

func (m *quesstionTest) TestGetInputError() {
	m.mockInput.On("getInput").Return("1", fmt.Errorf("input err"))
	i, err := getInput(m.exampleQuestion, m.mockInput)
	m.EqualError(err, "input err")
	m.Equal(-1, i)
}

func (m *quesstionTest) TestGetInputNotNumber() {
	m.mockInput.On("getInput").Return(fruit.Kiwi, nil).Once()
	m.mockInput.On("getInput").Return("1", nil).Once()
	i, err := getInput(m.exampleQuestion, m.mockInput)
	m.Nil(err)
	m.Equal(0, i)
}

func (m *quesstionTest) TestGetInputInvalidOption() {
	m.mockInput.On("getInput").Return("10000", nil).Once()
	m.mockInput.On("getInput").Return("1", nil).Once()
	i, err := getInput(m.exampleQuestion, m.mockInput)
	m.Nil(err)
	m.Equal(0, i)
}

func (m *quesstionTest) TestCheckAnswerCorrect() {
	correct := checkAnswer(m.exampleQuestion, 3)
	m.True(correct)
}

func (m *quesstionTest) TestCheckAnswerIncorrect() {
	correct := checkAnswer(m.exampleQuestion, 1)
	m.False(correct)
}

func (m *quesstionTest) TestRevealSuccessCorrect() {
	correct, err := revealSuccess(m.exampleQuestion, 3)
	m.True(correct)
	m.Nil(err)
}

func (m *quesstionTest) TestRevealSuccessIncorrect() {
	correct, err := revealSuccess(m.exampleQuestion, 1)
	m.False(correct)
	m.Nil(err)
}

func (m *quesstionTest) TestDisplayQuestion() {
	err := displayQuestion(m.exampleQuestion)
	m.Nil(err)
}
