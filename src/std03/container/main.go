package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	// HeapTest()

	// ListTest()

	RingTest()
}

func HeapTest() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Pop(h)
}

func ListTest() {
	li := list.New()
	li.PushFront(1)
	li.PushBack(2.98)
	ele := li.Back()
	li.InsertBefore("Hello", ele)
	fmt.Println(li.Len())
	for li.Len() != 0 {
		ele := li.Back()
		fmt.Print(ele.Value, " ")
		li.Remove(ele)
	}
}

func RingTest() {
	rings := ring.New(3)
	r := ring.New(2)

	fmt.Println(rings.Len())
	i := 0
	rings.Do(func(p interface{}) {
		rings.Value = i
		i++
		fmt.Print(p)
		rings = rings.Prev()
	})
	fmt.Println()
	rings = rings.Link(r)

	rings.Do(func(p interface{}) {
		fmt.Print(p, " ")
	})
	fmt.Println()

	fmt.Println(rings.Len())
	rings = rings.Unlink(2)
	rings.Do(func(p interface{}) {
		fmt.Print(p, " ")
	})
	fmt.Println()
	fmt.Println(rings.Len())
}
