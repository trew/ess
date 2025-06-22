package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var essCommand = &cobra.Command{
	Use:   "ess",
	Short: "Emergency Secret Sharing",
}

func Execute() {
	err := essCommand.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
