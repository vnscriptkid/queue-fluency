package main

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Define the workflow function with a delay.
func SimpleDelayedWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("SimpleDelayedWorkflow started")

	// Delay the workflow execution by 1 minute
	_ = workflow.Sleep(ctx, time.Minute)

	logger.Info("Executing workflow logic after delay")
	// Place your actual workflow logic here

	logger.Info("SimpleDelayedWorkflow completed")
	return nil
}

func main() {
	// Create Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		fmt.Println("Unable to create Temporal client", err)
		return
	}
	defer c.Close()

	// Create a worker to host workflow and activity functions
	w := worker.New(c, "simple-task-queue-delay", worker.Options{})

	// Register workflow and activity functions
	w.RegisterWorkflow(SimpleDelayedWorkflow)

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		fmt.Println("Unable to start worker", err)
	}
}
