package collectors

import (
	"context"
	"fmt"
	"time"
)

// DefaultTimeout is the per-collector timeout. Brew list can be slow
// on large setups, so 30s gives enough headroom.
const DefaultTimeout = 30 * time.Second

// CollectorResult is the output of a successful collection run.
type CollectorResult struct {
	Available bool          // true if the tool is installed
	Data      any           // tool-specific typed payload (e.g. *snapshot.HomebrewState)
	Warnings  []string      // non-fatal issues encountered during collection
	Duration  time.Duration // wall-clock time spent collecting
}

// Collector captures the current state of a single tool/package manager.
type Collector interface {
	// Name returns a unique identifier for this collector (e.g. "homebrew").
	Name() string

	// Available returns true if the tool is installed on this machine.
	// This should be a fast check (e.g. runner.Which).
	Available() bool

	// Collect gathers the current state of the tool.
	// Must respect context cancellation.
	Collect(ctx context.Context) (*CollectorResult, error)
}

// SafeCollect runs a collector with panic recovery and a per-collector timeout.
// It distinguishes "tool not installed" (returns result with Available=false)
// from "tool installed but collection failed" (returns error).
func SafeCollect(ctx context.Context, c Collector, timeout time.Duration) (*CollectorResult, error) {
	if !c.Available() {
		return &CollectorResult{Available: false}, nil
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	type outcome struct {
		res *CollectorResult
		err error
	}

	ch := make(chan outcome, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- outcome{nil, fmt.Errorf("collector %s panicked: %v", c.Name(), r)}
			}
		}()
		res, err := c.Collect(ctx)
		ch <- outcome{res, err}
	}()

	select {
	case o := <-ch:
		return o.res, o.err
	case <-ctx.Done():
		return nil, fmt.Errorf("collector %s: %w", c.Name(), ctx.Err())
	}
}
