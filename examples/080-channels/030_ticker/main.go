package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create context with cancel possibility
	ctx, cancel := context.WithCancel(context.Background())
	interval := 10 * time.Second

	// Run goroutine with ticker
	go func() {
		err := Run(ctx, interval)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}()

	// Run goroutine with ticker
	time.Sleep(5 * time.Second)

	// Cancel context to stop ticker
	cancel()

	// time to stop gourotine
	// in other case we may not see the ticker stopped message
	time.Sleep(10 * time.Second)

}
func Run(ctx context.Context, interval time.Duration) error {
	fmt.Println("starting worker")

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("running ticker at %s\n", interval.String())

			// do some job
			// check errors
		case <-ctx.Done():
			fmt.Printf("ticker stopped (%s)\n", ctx.Err().Error())

			return ctx.Err()
		}
	}

	return nil
}
