package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/trew/ess/ess"
	"golang.org/x/term"
	"os"
	"syscall"
)

var splitCommand = &cobra.Command{
	Use:   "split",
	Short: "Split a secret into N parts. K parts can be used to reconstruct the secret. Secret is prompted.",
	Run: func(cmd *cobra.Command, args []string) {
		partCount, err := cmd.Flags().GetInt("parts")
		if err != nil || partCount < 2 {
			fmt.Println("parts must be >= 2")
			os.Exit(1)
		}

		threshold, err := cmd.Flags().GetInt("threshold")
		if err != nil || threshold < 2 {
			fmt.Println("threshold must be >= 2")
			os.Exit(1)
		}

		fmt.Printf("Enter secret: ")
		pwd, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			fmt.Printf("Error reading password: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()

		parts, err := ess.Split(string(pwd), partCount, threshold)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for i, part := range parts {
			fmt.Printf("Part %d: %s\n", i+1, part)
		}
	},
}

func init() {
	essCommand.AddCommand(splitCommand)
	splitCommand.PersistentFlags().IntP("parts", "n", 0, "Number of parts to create")
	splitCommand.PersistentFlags().IntP("threshold", "k", 0, "Number of parts needed to reconstruct the secret")
}
