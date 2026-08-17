package waitnote

import (
	"testing"
	"time"
)

func TestBroadcastOrder(t *testing.T) {
	n := New()
	done := make(chan struct{})
	go func() {
		n.Wait()
		close(done)
	}()
	time.Sleep(50 * time.Millisecond)
	n.Signal()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("waiter not woken")
	}
}
