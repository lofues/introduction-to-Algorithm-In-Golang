package queue

import "errors"

type Queue struct {
	Size       int
	Data       []int
	Head, Tail int
	Total      int
}

func (q *Queue) Init(n int) {
	q.Data = make([]int, n)
	q.Size = 0
	q.Total = n
}

func (q *Queue) Length() int {
	return q.Tail - q.Head
}

func (q *Queue) Enqueue(num int) {
	q.Data[q.Tail] = num
	q.Tail++
}

func (q *Queue) Dequeue() (int, error) {
	if q.Length() == 0 {
		return 0, errors.New("without element")
	}
	res := q.Data[q.Head]
	q.Head++
	return res, nil
}

func (q *Queue) Top() (int, error) {
	if q.Length() == 0 {
		return 0, errors.New("without element")
	}
	return q.Data[q.Head], nil
}
