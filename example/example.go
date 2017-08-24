package example

import (
	"context"

	"github.com/tsavola/contextack"
)

type ackKey int

var (
	WaitDoneAck ackKey // Represents the termination of the Wait operation.
)

// Wait until ctx is cancelled, and acknowledge that it was respected.
func Wait(ctx context.Context) error {
	defer contextack.Ack(ctx, WaitDoneAck)
	<-ctx.Done()
	return ctx.Err()
}
