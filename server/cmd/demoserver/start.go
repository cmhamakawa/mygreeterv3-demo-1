package main

import (
	"os"
	"os/signal"
	"syscall"

	"dev.azure.com/service-hub-flg/service_hub_validation/_git/service_hub_validation_service.git/mygreeterv3/server/internal/demoserver"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the demoserver",
	Run:   start,
}

var options = demoserver.Options{}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntVar(&options.Port, "port", 50052, "the port to serve the demoserver on")
	startCmd.Flags().BoolVar(&options.JsonLog, "json-log", false, "The format of the log is json or user friendly key-value pairs")
}

func start(cmd *cobra.Command, args []string) {
	newServer := demoserver.NewServer()
	newServer.Init(options)
	newServer.Serve(options)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait indefinitely until a signal is received
	<-stop

	newServer.Cleanup()
}
