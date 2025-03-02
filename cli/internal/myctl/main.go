package myctl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/syunkitada/stadyapp/cli/internal/myctl/iam"
)

var rootCmd = &cobra.Command{
	Use:   "myctl",
	Short: "CLI for mycloudstack",
}

func init() {
	rootCmd.AddCommand(iam.RootCmd)
	// rootCmd.AddCommand(compute.RootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() {
	Execute()
}
