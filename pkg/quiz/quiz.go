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
		revealSuccess(q, optionIndex)
	}
	return nil
}
