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
		ID:        "user_registration_workflow",
		TaskQueue: "user-registration-task-queue",
	}

	userData := map[string]string{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, "UserRegistrationWorkflow", userData)
	if err != nil {
		fmt.Println("Unable to execute workflow", err)
		return
	}

	fmt.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
