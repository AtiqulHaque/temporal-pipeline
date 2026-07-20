package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/atiqulhaque/temporal-example/go-worker/internal/config"
	"github.com/atiqulhaque/temporal-example/go-worker/workflows"
	"go.temporal.io/sdk/client"
)

func main() {
	workflowID := flag.String("workflow-id", "", "workflow ID to signal (required)")
	runID := flag.String("run-id", "", "run ID (optional; omit if only one run is active)")
	approved := flag.Bool("approve", true, "approval value sent with the signal")
	signalName := flag.String("signal", workflows.ApproveSignal, "signal name")
	address := flag.String("address", config.TemporalAddress(), "Temporal server address")
	flag.Parse()

	if *workflowID == "" {
		fmt.Fprintln(os.Stderr, "error: --workflow-id is required")
		flag.Usage()
		os.Exit(2)
	}

	c, err := client.Dial(client.Options{HostPort: *address})
	if err != nil {
		log.Fatalf("unable to create Temporal client: %v", err)
	}
	defer c.Close()

	if err := c.SignalWorkflow(
		context.Background(),
		*workflowID,
		*runID,
		*signalName,
		*approved,
	); err != nil {
		log.Fatalf("unable to signal workflow: %v", err)
	}

	fmt.Printf("Sent signal %q to workflow %q (approve=%v)\n", *signalName, *workflowID, *approved)
}
