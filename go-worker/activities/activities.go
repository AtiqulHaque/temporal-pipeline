package activities

import (
	"context"
	"fmt"
	"log"
)

// GetGreeting returns a greeting for the given name (Go activity).
func GetGreeting(ctx context.Context, name string) (string, error) {
	log.Printf("[Go] GetGreeting activity started for: %s", name)
	return fmt.Sprintf("Hello, %s", name), nil
}


// HandleFailure runs when the workflow hits an error or rejection.
func HandleFailure(ctx context.Context, reason string) (string, error) {
	log.Printf("[Go] HandleFailure activity started: %s", reason)
	return fmt.Sprintf("Failure handled: %s", reason), nil
}

func LogResult(ctx context.Context, message string) error {
	log.Printf("[Go] Workflow completed with message: %s", message)
	return nil
}


// SearchUserMessage searches for a user message (Go activity).
func SearchUserMessage(ctx context.Context, message string) (string, error) {
	log.Printf("[Go] SearchUserMessage activity started for: %s", message)
	return fmt.Sprintf("User search message: %s", message), nil
}
