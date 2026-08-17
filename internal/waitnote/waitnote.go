package waitnote

import "sync"

type Note struct {
	mu sync.Mutex
	c  *sync.Cond
	ok bool
}

func New() *Note {
	n := &Note{}
	n.c = sync.NewCond(&n.mu)
	return n
}

func (n *Note) Wait() {
	n.mu.Lock()
	for !n.ok {
		n.c.Wait()
	}
	n.mu.Unlock()
}

func (n *Note) Signal() {
	n.mu.Lock()
	n.ok = true
	n.c.Broadcast()
	n.mu.Unlock()
}
