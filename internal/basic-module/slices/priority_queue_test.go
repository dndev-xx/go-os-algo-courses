//nolint:revive,errcheck,depguard,gosec,goconst
package slices_test

import (
	"testing"

	"github.com/dndev-xx/go-os-algo-courses/internal/basic-module/slices"
)

func TestEnqueue(t *testing.T) {
	tests := []struct {
		name        string
		maxPriority int
		enqueues    []struct {
			priority int
			value    int
		}
		wantSize int
		wantMinP int
		wantErr  bool
	}{
		{
			name:        "enqueue single item",
			maxPriority: 3,
			enqueues: []struct {
				priority int
				value    int
			}{
				{priority: 1, value: 100},
			},
			wantSize: 1,
			wantMinP: 1,
			wantErr:  false,
		},
		{
			name:        "enqueue multiple items same priority",
			maxPriority: 3,
			enqueues: []struct {
				priority int
				value    int
			}{
				{priority: 2, value: 200},
				{priority: 2, value: 201},
				{priority: 2, value: 202},
			},
			wantSize: 3,
			wantMinP: 2,
			wantErr:  false,
		},
		{
			name:        "enqueue multiple items different priorities",
			maxPriority: 3,
			enqueues: []struct {
				priority int
				value    int
			}{
				{priority: 3, value: 300},
				{priority: 1, value: 100},
				{priority: 2, value: 200},
				{priority: 0, value: 0},
			},
			wantSize: 4,
			wantMinP: 0,
			wantErr:  false,
		},
		{
			name:        "enqueue with invalid priority - too high",
			maxPriority: 3,
			enqueues: []struct {
				priority int
				value    int
			}{
				{priority: 4, value: 400},
			},
			wantSize: 0,
			wantMinP: 4,
			wantErr:  true,
		},
		{
			name:        "enqueue with invalid priority - negative",
			maxPriority: 3,
			enqueues: []struct {
				priority int
				value    int
			}{
				{priority: -1, value: -100},
			},
			wantSize: 0,
			wantMinP: 4,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := slices.NewPQ[int](tt.maxPriority)

			var lastErr error
			for _, e := range tt.enqueues {
				err := pq.Enqueue(e.priority, e.value)
				if err != nil {
					lastErr = err
				}
			}

			if tt.wantErr {
				if lastErr == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if lastErr != nil {
					t.Errorf("unexpected error: %v", lastErr)
				}
			}

			if pq.Len() != tt.wantSize {
				t.Errorf("size = %v, want %v", pq.Len(), tt.wantSize)
			}
			if !tt.wantErr && pq.MinP() != tt.wantMinP {
				t.Errorf("minP = %v, want %v", pq.MinP(), tt.wantMinP)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		name        string
		maxPriority int
		setup       func(pq *slices.PriorityQueue[int])
		wantItems   []int
		wantSize    int
		wantMinP    int
		wantErr     bool
	}{
		{
			name:        "dequeue from empty queue",
			maxPriority: 3,
			setup:       func(pq *slices.PriorityQueue[int]) {},
			wantItems:   []int{},
			wantSize:    0,
			wantMinP:    4,
			wantErr:     true,
		},
		{
			name:        "dequeue single item",
			maxPriority: 3,
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(1, 100)
			},
			wantItems: []int{100},
			wantSize:  0,
			wantMinP:  4,
			wantErr:   false,
		},
		{
			name:        "dequeue multiple items same priority (FIFO order)",
			maxPriority: 3,
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(1, 100)
				pq.Enqueue(1, 101)
				pq.Enqueue(1, 102)
			},
			wantItems: []int{100, 101, 102},
			wantSize:  0,
			wantMinP:  4,
			wantErr:   false,
		},
		{
			name:        "dequeue items with different priorities (priority order)",
			maxPriority: 3,
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(2, 200)
				pq.Enqueue(0, 0)
				pq.Enqueue(1, 100)
				pq.Enqueue(3, 300)
			},
			wantItems: []int{0, 100, 200, 300},
			wantSize:  0,
			wantMinP:  4,
			wantErr:   false,
		},
		{
			name:        "dequeue partially",
			maxPriority: 3,
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(0, 0)
				pq.Enqueue(0, 1)
				pq.Enqueue(1, 100)
				pq.Enqueue(1, 101)
			},
			wantItems: []int{0, 1, 100},
			wantSize:  1, // остался 101
			wantMinP:  1,
			wantErr:   false,
		},
		{
			name:        "dequeue with gaps in priorities",
			maxPriority: 5,
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(2, 200)
				pq.Enqueue(4, 400)
				pq.Enqueue(0, 0)
			},
			wantItems: []int{0, 200, 400},
			wantSize:  0,
			wantMinP:  6,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := slices.NewPQ[int](tt.maxPriority)
			tt.setup(pq)

			var gotItems []int
			for i := 0; i < len(tt.wantItems); i++ {
				item, err := pq.Dequeue()
				if err != nil {
					t.Errorf("unexpected error on dequeue %d: %v", i, err)
				}
				gotItems = append(gotItems, item)
			}

			for i, want := range tt.wantItems {
				if i < len(gotItems) && gotItems[i] != want {
					t.Errorf("item %d = %v, want %v", i, gotItems[i], want)
				}
			}

			if len(gotItems) != len(tt.wantItems) {
				t.Errorf("got %d items, want %d", len(gotItems), len(tt.wantItems))
			}

			if pq.Len() != tt.wantSize {
				t.Errorf("final size = %v, want %v", pq.Len(), tt.wantSize)
			}
			if pq.MinP() != tt.wantMinP {
				t.Errorf("final minP = %v, want %v", pq.MinP(), tt.wantMinP)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(pq *slices.PriorityQueue[int])
		wantEmpty bool
	}{
		{
			name:      "new queue is empty",
			setup:     func(pq *slices.PriorityQueue[int]) {},
			wantEmpty: true,
		},
		{
			name: "queue with items is not empty",
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(0, 100)
			},
			wantEmpty: false,
		},
		{
			name: "queue after enqueue and dequeue is empty",
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(0, 100)
				pq.Dequeue()
			},
			wantEmpty: true,
		},
		{
			name: "queue with multiple items after partial dequeue",
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(0, 100)
				pq.Enqueue(0, 101)
				pq.Dequeue()
			},
			wantEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := slices.NewPQ[int](5)
			tt.setup(pq)

			if got := pq.IsEmpty(); got != tt.wantEmpty {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.wantEmpty)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(pq *slices.PriorityQueue[int])
		wantPart string
	}{
		{
			name:     "empty queue string",
			setup:    func(pq *slices.PriorityQueue[int]) {},
			wantPart: "PriorityQueue: []",
		},
		{
			name: "queue with items string",
			setup: func(pq *slices.PriorityQueue[int]) {
				pq.Enqueue(1, 100)
				pq.Enqueue(0, 0)
			},
			wantPart: "PriorityQueue (size=2, minPriority=0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := slices.NewPQ[int](3)
			tt.setup(pq)

			got := pq.String()
			if tt.wantPart != "PriorityQueue: []" && got == "PriorityQueue: []" {
				t.Errorf("String() = %v, want contains %v", got, tt.wantPart)
			}
			if tt.wantPart == "PriorityQueue: []" && got != tt.wantPart {
				t.Errorf("String() = %v, want %v", got, tt.wantPart)
			}
		})
	}
}

func TestPriorityQueueWithStrings(t *testing.T) {
	pq := slices.NewPQ[string](2)

	err := pq.Enqueue(1, "middle")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = pq.Enqueue(0, "high")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = pq.Enqueue(2, "low")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	want := []string{"high", "middle", "low"}
	for i, w := range want {
		item, err := pq.Dequeue()
		if err != nil {
			t.Errorf("dequeue %d: unexpected error: %v", i, err)
		}
		if item != w {
			t.Errorf("dequeue %d = %v, want %v", i, item, w)
		}
	}

	if !pq.IsEmpty() {
		t.Error("queue should be empty")
	}
}

func TestPriorityQueueWithStructs(t *testing.T) {
	type Task struct {
		ID   int
		Name string
	}

	pq := slices.NewPQ[Task](1)

	task1 := Task{ID: 1, Name: "Task 1"}
	task2 := Task{ID: 2, Name: "Task 2"}

	err := pq.Enqueue(0, task1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = pq.Enqueue(1, task2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	item, err := pq.Dequeue()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if item.ID != 1 || item.Name != "Task 1" {
		t.Errorf("got %v, want Task 1", item)
	}

	item, err = pq.Dequeue()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if item.ID != 2 || item.Name != "Task 2" {
		t.Errorf("got %v, want Task 2", item)
	}
}
