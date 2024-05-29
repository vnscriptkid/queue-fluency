package main

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

func main() {
	// Create Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		fmt.Println("Unable to create Temporal client", err)
		return
	}
	defer c.Close()

	// Start SimpleWorkflowError execution
	workflowOptions := client.StartWorkflowOptions{
		ID:        "simple_workflow_error" + "_" + "try123",
		TaskQueue: "simple-task-queue-error",
		// If workflow is not completed within 10 minutes, timeout and terminate
		WorkflowExecutionTimeout: time.Minute * 3,
		WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING,
		RetryPolicy: &temporal.RetryPolicy{
			BackoffCoefficient: 1.0,
			MaximumInterval:    10 * time.Minute,
			InitialInterval:    3 * time.Second,
			MaximumAttempts:    20,
		},
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, "SimpleWorkflowError")
	if err != nil {
		fmt.Println("Unable to execute workflow", err)
		return
	}

	fmt.Println("Started SimpleWorkflowError", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
