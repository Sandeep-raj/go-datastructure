package utils

import "errors"

type Queue struct {
	Count    int
	Elements []interface{}
}

func InitQueue() *Queue {
	return &Queue{
		Count:    0,
		Elements: make([]interface{}, 0),
	}
}

func (q *Queue) Add(val interface{}) {
	q.Elements = append(q.Elements, val)
	q.Count++
}

func (q *Queue) Remove() (interface{}, error) {
	if q.Count == 0 {
		return nil, errors.New("queue is empty")
	}

	currVal := q.Elements[0]
	q.Elements = q.Elements[1:]
	q.Count--

	return currVal, nil
}

func (q *Queue) Size() int {
	return q.Count
}

func (q *Queue) Peek() (interface{}, error) {
	if q.Count == 0 {
		return nil, errors.New("queue is empty")
	}

	currVal := q.Elements[0]
	return currVal, nil
}
