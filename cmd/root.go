package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xm",
	Short: "xm",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("err: %v", err)
	}
}
