package quiz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:generate go run github.com/vektra/mockery/cmd/mockery -name inputter -inpkg --filename inputter_mock.go
type inputter interface {
	getInput() (string, error)
}

type keyboard struct{}

func (*keyboard) getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput, err
}

func validateInput(userInput string, optionCount int) (intInput int, valid bool, err error) {
	intInput, err = strconv.Atoi(userInput)
	if err != nil {
		_, err := fmt.Println("Please enter a number")
		return -1, false, err
	}
	if intInput > optionCount || intInput < 0 {
		_, err := fmt.Printf("Please enter a option between 1 and %d\n", optionCount)
		if err != nil {
			return -1, false, err
		}
		return -1, false, nil
	}
	return intInput, true, nil
}

func getInput(q question, i inputter) (int, error) {
	_, err := fmt.Print("What is your answer?: ")
	if err != nil {
		return -1, err
	}

	var intInput int
	valid := false
	optionCount := len(q.Options)
	for !valid {
		userInput, err := i.getInput()
		if err != nil {
			log.Println("err")
			return -1, err
		}
		intInput, valid, err = validateInput(userInput, optionCount)
		if err != nil {
			return -1, err
		}
	}
	return intInput - 1, nil
}

type question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

func checkAnswer(q question, answer int) bool {
	return q.Options[answer] == q.Answer
}

func revealSuccess(q question, answer int) (bool, error) {
	if checkAnswer(q, answer) {
		if _, err := fmt.Println("CORRECT! WELL DONE!"); err != nil {
			return false, err
		}
		return true, nil
	}
	_, err := fmt.Println("INCORRECT")
	return false, err
}

func displayQuestion(q question) error {
	if _, err := fmt.Printf("Q) %s \n", q.Question); err != nil {
		return err
	}
	for i, o := range q.Options {
		if _, err := fmt.Printf("\t%d) %s\n", i+1, o); err != nil {
			return err
		}
	}
	return nil
}
