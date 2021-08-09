
import (
	"container/heap"
	"fmt"
)
func minStoneSum(piles []int, k int) int {
    sortedarr := &IntHeap{}
    for _,i := range piles{
        heap.Push(sortedarr,i)
    }
    for i := 0 ; i < k ; i ++{
        g := heap.Pop(sortedarr).(int)
        g = g- int(g/2)
        heap.Push(sortedarr,g)
    }
    sum := 0 
    for sortedarr.Len() > 0{
        sum += heap.Pop(sortedarr).(int)
    }
    return sum
}


// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

