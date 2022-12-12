package utils

import "container/heap"

//
// PriorityQueue has been stolen from https://pkg.go.dev/container/heap
//

//
// HOWTO
//
// 1. Define struct which implements PQItem interface:
//		type Fruit struct {
//			name     string
//			priority int
//		}
//		func (i Fruit) LessThan(j utils.PQItem) bool {
//			return i.priority < j.(Fruit).priority
//		}
//
// 2. Create priority queue:
//		pq := utils.NewPriorityQueue[Fruit]()
//
// 3. Push new item:
//		pq.Push(&Fruit{name: "apple", priority: 3})
//
// 4. Pop from queue:
//		item := pq.Pop()
//
// 5. Success!
//

// An Item is something we manage in a priority queue.
type PQItem interface {
	LessThan(PQItem) bool // Defines the order of items in the priority queue (the smallest item comes first)
}

type PQ[T PQItem] struct {
	data heap_items[T]
}

func NewPq[T PQItem]() *PQ[T] {
	pq := &PQ[T]{
		data: make(heap_items[T], 0),
	}
	heap.Init(&pq.data)
	return pq
}

func (pq *PQ[T]) Push(item *T) {
	heap.Push(&pq.data, item)
}

func (pq *PQ[T]) Pop() *T {
	return heap.Pop(&pq.data).(*T)
}

func (pq *PQ[T]) Empty() bool {
	return pq.data.Len() == 0
}

// ------------------------------------- Internal priority queue realization -------------------------------------------

// An heap_item is something we manage in a priority queue.
type heap_item[T PQItem] struct {
	value *T  // user data
	index int // index of the item in the heap
}

type heap_items[T PQItem] []*heap_item[T]

func (pq heap_items[T]) Len() int {
	return len(pq)
}

func (pq heap_items[T]) Less(i, j int) bool {
	v1 := pq[i].value
	v2 := pq[j].value
	return (*v1).LessThan(*v2)
}

func (pq heap_items[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *heap_items[T]) Push(x any) {
	n := len(*pq)
	value := x.(*T)
	item := &heap_item[T]{value: value, index: n}
	*pq = append(*pq, item)
}

func (pq *heap_items[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item.value
}

// ------------------------------------- Internal priority queue realization -------------------------------------------
