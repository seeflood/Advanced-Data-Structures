package deque

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	dq := NewLinkedList()
	assertTrue(dq.IsEmpty(), t)
	dq.AddFirst(1)
	dq.AddFirst(0)
	assertTrue(dq.Size() == 2, t)
	assertTrue(dq.PeekFirst() == 0, t)
	assertTrue(dq.PeekLast() == 1, t)
	assertTrue(dq.PollFirst() == 0, t)
	assertTrue(dq.Size() == 1, t)
	assertTrue(!dq.IsEmpty(), t)
	assertTrue(!dq.AddFirst(nil), t)
	assertTrue(dq.AddFirst(666), t)
	assertTrue(dq.Size() == 2, t)
	assertTrue(dq.PeekLast() == 1, t)
}
