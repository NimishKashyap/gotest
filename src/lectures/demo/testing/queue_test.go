package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue count: %v, want %v", len(q.items), i)
		}

		if !q.Append(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}

	if q.Append(4) {
		t.Errorf("Should not be able to add to full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()

		if !ok {
			t.Errorf("Should not be able to get item from queue")
		}

		if item != i {
			t.Errorf("Wrong Order: %v, want %v", item, i)
		}
	}

	item, ok := q.Next()

	if ok {
		t.Errorf("Should be any more items in queu, got, %v", item)
	}
}
