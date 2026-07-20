package workflows

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

const GoTaskQueue = "go-task-queue"
const NestJSTaskQueue = "nestjs-task-queue"

const GetGreetingActivity = "getGreeting"
const FormatMessageActivity = "formatMessage"
const CreateInvoiceActivity = "createInvoice"
const LogResultActivity = "logResult"
const SearchUserMessageActivity = "searchUserMessage"
const HandleFailureActivity = "handleFailure"

const ApproveSignal = "approve"
const approveTimeout = 5 * time.Minute

// CreateInvoiceInput is the payload for the NestJS createInvoice activity.
type CreateInvoiceInput struct {
	CustomerName  string  `json:"customerName"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency,omitempty"`
	WorkflowID    string  `json:"workflowId,omitempty"`
	InvoiceNumber string  `json:"invoiceNumber,omitempty"`
}

// CreateInvoiceResult is returned by the NestJS createInvoice activity.
type CreateInvoiceResult struct {
	ID            string `json:"id"`
	InvoiceNumber string `json:"invoiceNumber"`
	CustomerName  string `json:"customerName"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
	WorkflowID    string `json:"workflowId"`
}

// GreetingWorkflow orchestrates activities across Go and NestJS workers.
// After getGreeting it waits for an "approve" signal before continuing.
func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)

	goAO := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		TaskQueue:           GoTaskQueue,
	}
	ctx = workflow.WithActivityOptions(ctx, goAO)

	var greeting string
	if err := workflow.ExecuteActivity(ctx, GetGreetingActivity, name).Get(ctx, &greeting); err != nil {
		return runFailureHandler(ctx, fmt.Sprintf("getGreeting failed: %v", err))
	}

	logger.Info("Waiting for approve signal", "signal", ApproveSignal, "timeout", approveTimeout)

	approved, err := waitForApproval(ctx)
	if err != nil {
		return runFailureHandler(ctx, fmt.Sprintf("approve signal error: %v", err))
	}
	if !approved {
		return runFailureHandler(ctx, fmt.Sprintf("workflow rejected or timed out waiting for %q signal", ApproveSignal))
	}

	nestjsAO := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		TaskQueue:           NestJSTaskQueue,
	}
	nestjsCtx := workflow.WithActivityOptions(ctx, nestjsAO)

	var message string
	if err := workflow.ExecuteActivity(nestjsCtx, FormatMessageActivity, greeting).Get(ctx, &message); err != nil {
		return runFailureHandler(ctx, fmt.Sprintf("formatMessage failed: %v", err))
	}

	// Persist invoice only after approve (this block runs only when approved == true).
	var invoice CreateInvoiceResult
	if err := workflow.ExecuteActivity(nestjsCtx, CreateInvoiceActivity, CreateInvoiceInput{
		CustomerName: name,
		Amount:       100.00,
		Currency:     "USD",
		WorkflowID:   workflow.GetInfo(ctx).WorkflowExecution.ID,
	}).Get(ctx, &invoice); err != nil {
		return runFailureHandler(ctx, fmt.Sprintf("createInvoice failed: %v", err))
	}
	logger.Info("Invoice created", "invoiceNumber", invoice.InvoiceNumber, "invoiceId", invoice.ID)

	if err := workflow.ExecuteActivity(ctx, LogResultActivity, message).Get(ctx, nil); err != nil {
		return runFailureHandler(ctx, fmt.Sprintf("logResult failed: %v", err))
	}

	var userMessage string
	if err := workflow.ExecuteActivity(ctx, SearchUserMessageActivity, message).Get(ctx, &userMessage); err != nil {
		return runFailureHandler(ctx, fmt.Sprintf("searchUserMessage failed: %v", err))
	}

	return userMessage, nil
}

func waitForApproval(ctx workflow.Context) (bool, error) {
	logger := workflow.GetLogger(ctx)
	approveCh := workflow.GetSignalChannel(ctx, ApproveSignal)

	var approved bool
	timer := workflow.NewTimer(ctx, approveTimeout)

	selector := workflow.NewSelector(ctx)
	selector.AddReceive(approveCh, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &approved)
		logger.Info("Received approve signal", "approved", approved)
	})
	selector.AddFuture(timer, func(f workflow.Future) {
		logger.Warn("Timed out waiting for approve signal")
		approved = false
	})
	selector.Select(ctx)

	return approved, nil
}

func runFailureHandler(ctx workflow.Context, reason string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Warn("Running failure handler activity", "reason", reason)

	var note string
	if err := workflow.ExecuteActivity(ctx, HandleFailureActivity, reason).Get(ctx, &note); err != nil {
		return "", fmt.Errorf("failure handler activity failed: %w (original: %s)", err, reason)
	}

	return note, nil
}
