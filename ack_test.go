package contextack_test

import (
	"context"
	"testing"

	"github.com/tsavola/contextack"
	"github.com/tsavola/contextack/example"
)

func Test(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, ack1 := contextack.WithAck(ctx, example.WaitDoneAck)
	ctx, ack2 := contextack.WithAck(ctx, example.WaitDoneAck)
	go example.Wait(ctx)

	cancel()
	<-ack1
	<-ack2

	// Second call should panic due to double ack.
	defer func() {
		if x := recover(); x != nil {
			t.Logf("%#v", x)
		} else {
			t.Fail()
		}
	}()
	example.Wait(ctx)
}
