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

func (k *keyboard) getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput, err
}

func getInput(q question, i inputter) (int, error) {
	fmt.Print("What is your answer?: ")
	var intInput int
	valid := false
	optionCount := len(q.Options)
	for !valid {
		userInput, err := i.getInput()
		if err != nil {
			log.Println("err")
			return -1, err
		}
		intInput, err = strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			if intInput < optionCount+1 && intInput > 0 {
				valid = true
			} else {
				fmt.Printf("Please enter a option between 1 and %d\n", optionCount)
			}

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

func revealSuccess(q question, answer int) bool {
	if checkAnswer(q, answer) {
		fmt.Println("CORRECT! WELL DONE!")
		return true
	}
	fmt.Println("INCORRECT")
	return false
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
