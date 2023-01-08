package main

import (
	"context"

	"github.com/julientant/example-go-logger-ctx/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	ctx = logger.WithLogger(ctx, logrus.StandardLogger())
	logger.FromContext(ctx).Info("hello world")
	someDBOperation(ctx)
	logger.FromContext(ctx).Info("bye")
}

func someDBOperation(ctx context.Context) {
	ctx = logger.WithNonSpillingField(ctx, "enrich_logger", "yes")
	ctx = logger.WithField(ctx, "enrich_context", "yes")
	logger.FromContext(ctx).Info("some db operation")
	needAssistance(ctx)
}

func needAssistance(ctx context.Context) {
	logger.FromContext(ctx).Info("need assistance")
}
