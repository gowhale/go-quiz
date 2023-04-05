// Package quiz is responsible for generating and playing a quiz
package quiz

func playQuiz(questions []question, i inputter) error {
	for _, q := range questions {
		if err := displayQuestion(q); err != nil {
			return err
		}
		optionIndex, err := getInput(q, i)
		if err != nil {
			return err
		}
		if _, err := revealSuccess(q, optionIndex); err != nil {
			return err
		}
	}
	return nil
}
