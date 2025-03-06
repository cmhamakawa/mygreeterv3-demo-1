package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "demoserver",
	Short: "Demoserver is a standalone component for handling SayHello requests",
	Long:  `Demoserver is a standalone component that handles SayHello requests and can be run independently.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
