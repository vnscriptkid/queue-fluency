package main

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		fmt.Println("Unable to create Temporal client", err)
		return
	}
	defer c.Close()

	// Start workflow execution
	workflowOptions := client.StartWorkflowOptions{
		ID:        "simple_workflow_delay",
		TaskQueue: "simple-task-queue-delay",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, "SimpleDelayedWorkflow")
	if err != nil {
		fmt.Println("Unable to execute workflow", err)
		return
	}

	fmt.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
