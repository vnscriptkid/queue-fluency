package main

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Define the workflow function.
func SimpleWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("SimpleWorkflow started")

	// Sleep for 1 second
	_ = workflow.Sleep(ctx, time.Second*1)

	logger.Info("SimpleWorkflow completed")
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
	w := worker.New(c, "simple-task-queue", worker.Options{})

	// Register workflow and activity functions
	w.RegisterWorkflow(SimpleWorkflow)

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		fmt.Println("Unable to start worker", err)
	}
}
