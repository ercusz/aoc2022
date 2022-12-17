package utils

type Queue[T any] []T

func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(*q) == 0 {
		var v T
		return v, false
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, true
}
