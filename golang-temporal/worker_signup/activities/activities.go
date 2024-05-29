package activities

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
)

// Activity to validate user data
func ValidateUserDataActivity(ctx context.Context, userData map[string]string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Validating user data", "UserData", userData)

	// Simulate validation delay
	time.Sleep(1 * time.Second)

	// Add validation logic here (e.g., check if email is valid)
	if userData["email"] == "" {
		return fmt.Errorf("invalid email")
	}

	logger.Info("User data validated successfully")
	return nil
}

// Activity to create user account
func CreateUserAccountActivity(ctx context.Context, userData map[string]string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Creating user account", "UserData", userData)

	// Simulate account creation delay
	time.Sleep(2 * time.Second)

	// Add account creation logic here (e.g., save user data to database)
	logger.Info("User account created successfully")
	return nil
}

// Activity to send welcome email
func SendWelcomeEmailActivity(ctx context.Context, email string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Sending welcome email", "Email", email)

	// Simulate email sending delay
	time.Sleep(2 * time.Second)

	logger.Info("Welcome email sent successfully")
	return nil
}
