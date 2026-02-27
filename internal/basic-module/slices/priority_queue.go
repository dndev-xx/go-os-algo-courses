//nolint:revive
package slices

import (
	"errors"
	"fmt"
	"strings"
)

type PriorityQueue[T any] struct {
	queues [][]T // index = priority
	minP   int
	maxP   int
	size   int
}

func NewPQ[T any](maxPriority int) *PriorityQueue[T] {
	if maxPriority < 0 {
		maxPriority = 0
	}
	queues := make([][]T, maxPriority+1)
	for i := range queues {
		queues[i] = make([]T, 0)
	}
	return &PriorityQueue[T]{
		queues: queues,
		maxP:   maxPriority,
		minP:   maxPriority + 1,
		size:   0,
	}
}

func (pq *PriorityQueue[T]) Enqueue(priority int, value T) error {
	if priority < 0 || priority > pq.maxP {
		return fmt.Errorf("priority out of range: %d (min: 0, max: %d)", priority, pq.maxP)
	}
	pq.queues[priority] = append(pq.queues[priority], value)
	if priority < pq.minP {
		pq.minP = priority
	}
	pq.size++
	return nil
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	var zero T
	if pq.IsEmpty() {
		return zero, errors.New("priority queues is empty")
	}

	for pq.minP <= pq.maxP {
		queue := pq.queues[pq.minP]
		if len(queue) > 0 {
			item := queue[0]
			if len(queue) == 1 {
				pq.queues[pq.minP] = make([]T, 0)
			} else {
				pq.queues[pq.minP] = queue[1:]
			}

			pq.size--

			if len(pq.queues[pq.minP]) == 0 {
				pq.updateMinPriority()
			}

			return item, nil
		}
	}

	return zero, nil
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.size
}

func (pq *PriorityQueue[T]) MinP() int {
	return pq.minP
}

func (pq *PriorityQueue[T]) String() string {
	if pq.IsEmpty() {
		return "PriorityQueue: []"
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("PriorityQueue (size=%d, minPriority=%d):\n", pq.size, pq.minP))
	for i := pq.minP; i <= pq.maxP; i++ {
		if len(pq.queues[i]) > 0 {
			sb.WriteString(fmt.Sprintf("  Priority %d (%d items): %v\n", i, len(pq.queues[i]), pq.queues[i]))
		}
	}
	return sb.String()
}

func (pq *PriorityQueue[T]) updateMinPriority() {
	for i := pq.minP; i <= pq.maxP; i++ {
		if len(pq.queues[i]) > 0 {
			pq.minP = i
			return
		}
	}
	pq.minP = pq.maxP + 1
}
