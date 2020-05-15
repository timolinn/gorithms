package datastructures

import (
	"errors"
)

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	var item interface{}

	if q.IsEmpty() {
		return nil, errors.New("Cannot remove from empty queue")
	}
	item = q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Peek returns the top element
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		panic("empty queue")
	}
	return q.items[0]
}

// IsEmpty returns false if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}
