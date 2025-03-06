package main

import (
	"context"
	"io"
	"os"

	log "log/slog"

	demoserver "dev.azure.com/service-hub-flg/service_hub_validation/_git/service_hub_validation_service.git/mygreeterv3/server/internal/demoserver"
	"dev.azure.com/service-hub-flg/service_hub_validation/_git/service_hub_validation_service.git/mygreeterv3/server/internal/logattrs"
	"github.com/Azure/aks-middleware/grpc/server/ctxlogger"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the demoserver service",
	Run:   start,
}

var options = demoserver.Options{}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntVar(&options.Port, "port", 50053, "the port to serve the demoserver on")
	startCmd.Flags().BoolVar(&options.JsonLog, "json-log", false, "The format of the log is json or user friendly key-value pairs")
	startCmd.Flags().StringVar(&options.RemoteAddr, "remote-addr", "", "the remote server's address for this demoserver to connect to")
}

var output io.Writer = os.Stdout

func start(cmd *cobra.Command, args []string) {
	logger := log.New(log.NewTextHandler(output, nil).WithAttrs(logattrs.GetAttrs()))
	if options.JsonLog {
		logger = log.New(log.NewJSONHandler(output, nil).WithAttrs(logattrs.GetAttrs()))
	}

	log.SetDefault(logger)
	ctx, cancel := context.WithCancel(context.Background())
	ctx = ctxlogger.WithLogger(ctx, logger)
	defer cancel()
	demoServer, err := demoserver.NewServer(ctx, options)
	if err != nil {
		logger.Error("Something went wrong starting demoserver: " + err.Error())
		os.Exit(1)
	}

	err = demoServer.Serve(ctx)
	if err != nil {
		logger.Error("Something went wrong running demoserver.")
	}
}
