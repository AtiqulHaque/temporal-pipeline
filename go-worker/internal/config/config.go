package config

import "os"

func EnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func TemporalAddress() string {
	return EnvOrDefault("TEMPORAL_ADDRESS", "localhost:7233")
}

// GoTaskQueue must match workflows.GoTaskQueue.
func GoTaskQueue() string {
	return EnvOrDefault("GO_TASK_QUEUE", EnvOrDefault("TASK_QUEUE", "go-task-queue"))
}

// NestJSTaskQueue must match workflows.NestJSTaskQueue.
func NestJSTaskQueue() string {
	return EnvOrDefault("NESTJS_TASK_QUEUE", "nestjs-task-queue")
}
