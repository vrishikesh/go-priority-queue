package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Name     string
	Priority int
}

type PriorityQueue []Item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Priority < (*pq)[j].Priority
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Swap(i, j int) {
	cp := *pq
	cp[i], cp[j] = cp[j], cp[i]
}

func main() {
	listItems := []Item{
		{Name: "Carrot", Priority: 30},
		{Name: "Potato", Priority: 45},
		{Name: "Rice", Priority: 100},
		{Name: "Spinach", Priority: 5},
	}
	var pq PriorityQueue
	// pq := make(PriorityQueue, len(listItems))
	// copy(pq, listItems)
	heap.Init(&pq)

	for _, item := range listItems {
		heap.Push(&pq, item)
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(Item)
		fmt.Println(item)
	}
}
