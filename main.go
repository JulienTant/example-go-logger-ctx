package main

import (
	"context"

	"github.com/julientant/example-go-logger-ctx/log"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	ctx = log.WithLogger(ctx, logrus.StandardLogger())
	// we could do without the line above to use the default logger

	log.FromContext(ctx).Info("hello from main, no fields")

	// we can enrich the context with fields
	ctx = log.WithField(ctx, "user_id", "user_123")
	log.FromContext(ctx).Info("hello from main, with user_id field")

	ctx = log.WithFields(ctx, logrus.Fields{
		"request_id": "request_123",
		"session_id": "session_123",
	})
	log.FromContext(ctx).Info("hello from main, with request_id and session_id fields")

	secondFunc(ctx)

	log.FromContext(ctx).Info("after second func, the context is still enriched")
}

func secondFunc(ctx context.Context) {
	ctx = log.WithField(ctx, "added_from_second_func", "yes")

	log.FromContext(ctx).Info("with our new field")
}

// Output:
// INFO[0000] hello from main, no fields
// INFO[0000] hello from main, with user_id field           user_id=user_123
// INFO[0000] hello from main, with request_id and session_id fields  request_id=request_123 session_id=session_123 user_id=user_123
// INFO[0000] with our new field                            added_from_second_func=yes request_id=request_123 session_id=session_123 user_id=user_123
// INFO[0000] after second func, the context is still enriched  added_from_second_func=yes request_id=request_123 session_id=session_123 user_id=user_123
