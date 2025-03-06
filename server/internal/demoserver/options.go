package demoserver

import (
	"github.com/spf13/cobra"
)

type Options struct {
	Port    int
	JsonLog bool
}

func ParseOptions(cmd *cobra.Command) Options {
	var options Options
	cmd.Flags().IntVar(&options.Port, "port", 50052, "the port to serve the demoserver on")
	cmd.Flags().BoolVar(&options.JsonLog, "json-log", false, "The format of the log is json or user friendly key-value pairs")
	return options
}
