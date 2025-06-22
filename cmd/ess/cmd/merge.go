package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/trew/ess/ess"
	"os"
)

var mergeCommand = &cobra.Command{
	Use:   "merge [parts...]",
	Short: "Merge a number of parts to reconstruct the secret",
	Run: func(cmd *cobra.Command, args []string) {
		merged, err := ess.Merge(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Secret: %s\n", merged)
	},
}

func init() {
	essCommand.AddCommand(mergeCommand)
}
