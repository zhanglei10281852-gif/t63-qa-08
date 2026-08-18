package health

import (
	"context"
	"testing"
	"time"
)

func TestCheckerStopsWhenParentRequestIsCancelled(t *testing.T) {
	started := make(chan struct{})
	checker := Checker{Timeout: 5 * time.Second, Checks: map[string]Check{
		"database": func(ctx context.Context) error {
			close(started)
			<-ctx.Done()
			return ctx.Err()
		},
	}}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { _, err := checker.Run(ctx); done <- err }()
	<-started
	cancel()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("health checker ignored parent cancellation")
	}
}
