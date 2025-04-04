package actions

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"
)

func HelloActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("HelloActivity", "name", name)
	return fmt.Sprintf("HelloActivity, %s!", name), nil
}

func ByeActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ByeActivity", "name", name)
	return fmt.Sprintf("ByeActivity, %s!", name), nil
}
