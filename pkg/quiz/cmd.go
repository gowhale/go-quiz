package quiz

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		questions, err := getQuestions(&local{})
		if err != nil {
			return err
		}
		return playQuiz(questions, &keyboard{})
	},
}

// Execute adds the quiz flag to the rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
