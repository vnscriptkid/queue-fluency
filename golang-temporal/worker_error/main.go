package main

import (
	"errors"
	"fmt"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

var ErrExample = errors.New("example error")

// Define the workflow function.
func SimpleWorkflowError(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info(">>1. SimpleWorkflowError started")

	// Simulate some logic that could potentially fail
	success, err := performLogic()
	if err != nil {
		logger.Error(">>2. Logic failed", "Error", err)
		// Handle error by retrying or compensating
		return err
	}

	if !success {
		err := ErrExample
		logger.Error(">>3. Logic returned false", "Error", err)
		// Handle error by retrying or compensating
		return err
	}

	logger.Info(">>4. Logic succeeded")

	// Simulate waiting for a certain duration
	workflow.Sleep(ctx, time.Second*5)

	logger.Info(">>5. SimpleWorkflowError completed successfully")
	return nil
}

// performLogic simulates some internal logic within the workflow
func performLogic() (bool, error) {
	// Here, we're simulating a failure scenario with a 50% chance of failure
	nowUnix := time.Now().Unix()

	val := nowUnix % 2

	if val == 0 {
		return false, fmt.Errorf("!!failed: now = %v, val = %v", nowUnix, val)
	}
	return true, nil
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
	w := worker.New(c, "simple-task-queue-error", worker.Options{})

	// Register workflow and activity functions
	w.RegisterWorkflow(SimpleWorkflowError)

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		fmt.Println("Unable to start worker", err)
	}
}
