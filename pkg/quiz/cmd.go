package quiz

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func checkAnswer(q Question, answer int) bool {
	return q.Options[answer] == q.Answer
}

func revealSuccess(q Question, answer int) {
	if checkAnswer(q, answer) {
		fmt.Println("CORRECT! WELL DONE!")
	} else {
		fmt.Println("INCORRECT")
	}
}

func playQuiz(questions []Question) {
	for _, q := range questions {
		displayQuestion(q)
		a, _ := getInput(q)
		revealSuccess(q, a)
	}
}

func getInput(q Question) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your answer?: ")
	var intInput int
	var err error
	valid := false
	optionCount := len(q.Options)
	for !valid {
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
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

func displayQuestion(q Question) {
	fmt.Printf("Q) %s \n", q.Question)
	for i, o := range q.Options {
		fmt.Printf("\t%d) %s\n", i+1, o)
	}
}

func getQuestions() (questions []Question, err error) {
	jsonFile, err := os.Open("metadata/1/questions.json")
	if err != nil {
		return []Question{}, err
	}

	defer func() {
		if err2 := jsonFile.Close(); err2 != nil {
			if err == nil {
				err = err2
			} else {
				err = fmt.Errorf("%w %w", err, err2)
			}
		}
	}()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(byteValue, &questions)
	return questions, err
}

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		questions, err := getQuestions()
		if err != nil {
			fmt.Println(err)
		}
		playQuiz(questions)
	},
}

type Question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
