// Package contextack extends context.Context with cancellation
// acknowledgement.
//
// API implementations should have context value keys for functions which
// support context.  A group of functions may share the same key if only one of
// them will be called with a given context (including its parent contexts).
// The functions should call Ack() with the key just before exiting.
//
// API clients should call WithAck() with the key which represents the function
// they are about to call (and possibly cancel).  When it's important to wait
// until the function has finished, the receive channel returned by WithAck can
// be used.
package contextack

import (
	"context"
)

// WithAck returns ctx or a copy of it, and a channel which will be closed when
// the operation represented by the key has finished.
func WithAck(ctx context.Context, key interface{}) (context.Context, <-chan struct{}) {
	var ack chan struct{}
	if x := ctx.Value(key); x != nil {
		ack = x.(chan struct{})
	} else {
		ack = make(chan struct{})
		ctx = context.WithValue(ctx, key, ack)
	}
	return ctx, ack
}

// Ack notifies all subscribers that the operation represented by the key has
// finished.  It's ok to call this even when there are no subscribers.
func Ack(ctx context.Context, key interface{}) {
	if x := ctx.Value(key); x != nil {
		ack := x.(chan struct{})
		close(ack)
	}
}
