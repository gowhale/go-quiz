package quiz

import (
	"fmt"
	"testing"

	fruit "github.com/gowhale/go-test-data/pkg/fruits"
	"github.com/stretchr/testify/suite"
)

const ()

type quizTest struct {
	suite.Suite

	exampleQuiz []question
	mockInput   *mockInputter
}

func (m *quizTest) SetupTest() {
	m.exampleQuiz = []question{
		{Question: fruit.Apple,
			Options: []string{
				fruit.Apricot,
				fruit.Avocado,
				fruit.Blueberry,
				fruit.Lemon,
			},
			Answer: fruit.Lemon,
		}}
	m.mockInput = new(mockInputter)
}

func TestQuizTest(t *testing.T) {
	suite.Run(t, new(quizTest))
}

func (m *quizTest) TestPlayQuizPass() {
	m.mockInput.On("getInput").Return("1", nil)
	err := playQuiz(m.exampleQuiz, m.mockInput)
	m.Nil(err)
}

func (m *quizTest) TestPlayQuizFail() {
	m.mockInput.On("getInput").Return("", fmt.Errorf("input err"))
	err := playQuiz(m.exampleQuiz, m.mockInput)
	m.EqualError(err, "input err")
}
