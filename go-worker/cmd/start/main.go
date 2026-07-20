package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/atiqulhaque/temporal-example/go-worker/internal/config"
	"github.com/atiqulhaque/temporal-example/go-worker/workflows"
	"go.temporal.io/sdk/client"
)

func main() {
	name := flag.String("name", "Temporal", "name passed to GreetingWorkflow")
	workflowID := flag.String("id", "", "workflow ID (default: greeting-workflow-<timestamp>)")
	taskQueue := flag.String("task-queue", config.GoTaskQueue(), "Temporal task queue for the workflow")
	address := flag.String("address", config.TemporalAddress(), "Temporal server address")
	wait := flag.Bool("wait", true, "wait for workflow to complete and print result")
	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "unexpected arguments: %v\n", flag.Args())
		flag.Usage()
		os.Exit(2)
	}

	if *workflowID == "" {
		*workflowID = "greeting-workflow-" + time.Now().Format("20060102-150405")
	}

	c, err := client.Dial(client.Options{HostPort: *address})
	
	if err != nil {
		log.Fatalf("unable to create Temporal client: %v", err)
	}
	defer c.Close()

	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        *workflowID,
		TaskQueue: *taskQueue,
	}, workflows.GreetingWorkflow, *name)

	if err != nil {
		log.Fatalf("unable to start workflow: %v", err)
	}

	fmt.Printf("Started workflow\n  WorkflowID: %s\n  RunID:      %s\n", we.GetID(), we.GetRunID())
	fmt.Printf("\nWorkflow waits for approve signal before continuing.\n")
	fmt.Printf("  go run ./cmd/signal --workflow-id %s --approve=true\n\n", we.GetID())

	if !*wait {
		return
	}

	var result string
	if err := we.Get(context.Background(), &result); err != nil {
		log.Fatalf("workflow failed: %v", err)
	}

	fmt.Printf("Result: %s\n", result)
}
