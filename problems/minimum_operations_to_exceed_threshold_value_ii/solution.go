
import "container/heap"
func minOperations(nums []int, k int) int {
    h := &IntHeap{}
    for _, i := range nums {
        heap.Push(h, i)
    }
    cnt := 0
    for h.Len() > 1 {
        a, b := heap.Pop(h).(int), heap.Pop(h).(int)
        if a >= k {
            return cnt
        }
        cnt ++
        heap.Push(h, min(a, b) * 2 + max(a, b))
    }
    return len(nums) - 1
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}