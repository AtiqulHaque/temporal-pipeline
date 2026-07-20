package main

import (
	"log"

	"github.com/atiqulhaque/temporal-example/go-worker/activities"
	"github.com/atiqulhaque/temporal-example/go-worker/internal/config"
	"github.com/atiqulhaque/temporal-example/go-worker/workflows"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{HostPort: config.TemporalAddress()})
	if err != nil {
		log.Fatalf("unable to create Temporal client: %v", err)
	}
	defer c.Close()

	taskQueue := config.GoTaskQueue()
	w := worker.New(c, taskQueue, worker.Options{})
	w.RegisterWorkflow(workflows.GreetingWorkflow)

	// Register with explicit names so they match the workflow and NestJS worker.
	w.RegisterActivityWithOptions(activities.GetGreeting, activity.RegisterOptions{Name: workflows.GetGreetingActivity})
	w.RegisterActivityWithOptions(activities.LogResult, activity.RegisterOptions{Name: workflows.LogResultActivity})
	w.RegisterActivityWithOptions(activities.SearchUserMessage, activity.RegisterOptions{Name: workflows.SearchUserMessageActivity})
	w.RegisterActivityWithOptions(activities.HandleFailure, activity.RegisterOptions{Name: workflows.HandleFailureActivity})

	log.Printf("Go worker listening on task queue %q", taskQueue)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("unable to start worker: %v", err)
	}
}
