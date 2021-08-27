import (
	"container/heap"
	"fmt"
)
func largestSumAfterKNegations(nums []int, k int) int {
    h := &IntHeap{}
    for _,i := range  nums{
        heap.Push(h,i)
    }
    for i := 0 ; i < k ; i++{
        temp := heap.Pop(h).(int)
        temp *= -1
        heap.Push(h,temp)
    }
    sum := 0 
    for h.Len() > 0{
        sum += heap.Pop(h).(int)
    }
    return sum
}


// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}