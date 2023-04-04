package quiz

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		questions, err := getQuestions(&local{})
		if err != nil {
			fmt.Println(err)
		}
		playQuiz(questions, &keyboard{})
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
