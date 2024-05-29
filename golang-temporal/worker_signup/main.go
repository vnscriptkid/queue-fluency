package main

import (
	"fmt"
	"time"

	"github.com/vnscriptkid/queue-fluency/golang-temporal/worker_signup/activities"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// Define the workflow function with multiple activities.
func UserRegistrationWorkflow(ctx workflow.Context, userData map[string]string) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("UserRegistrationWorkflow started")

	// Activity options
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Call ValidateUserDataActivity
	err := workflow.ExecuteActivity(ctx, activities.ValidateUserDataActivity, userData).Get(ctx, nil)
	if err != nil {
		logger.Error("ValidateUserDataActivity failed.", "Error", err)
		return err
	}

	// Call CreateUserAccountActivity
	err = workflow.ExecuteActivity(ctx, activities.CreateUserAccountActivity, userData).Get(ctx, nil)
	if err != nil {
		logger.Error("CreateUserAccountActivity failed.", "Error", err)
		return err
	}

	// Call SendWelcomeEmailActivity
	err = workflow.ExecuteActivity(ctx, activities.SendWelcomeEmailActivity, userData["email"]).Get(ctx, nil)
	if err != nil {
		logger.Error("SendWelcomeEmailActivity failed.", "Error", err)
		return err
	}

	logger.Info("UserRegistrationWorkflow completed")
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
	w := worker.New(c, "user-registration-task-queue", worker.Options{})

	// Register workflow and activity functions
	w.RegisterWorkflow(UserRegistrationWorkflow)
	w.RegisterActivity(activities.ValidateUserDataActivity)
	w.RegisterActivity(activities.CreateUserAccountActivity)
	w.RegisterActivity(activities.SendWelcomeEmailActivity)

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		fmt.Println("Unable to start worker", err)
	}
}
