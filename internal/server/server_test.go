package server

import (
	"context"
	"testing"
	"time"
)

func TestStartStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan error, 1)
	go func() {
		done <- NewServer(0).Start(ctx) // port 0: any free port
	}()

	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Start returned %v, want nil after cancel", err)
		}
	case <-time.After(shutdownTimeout + time.Second):
		t.Fatal("Start did not return after cancel")
	}
}
