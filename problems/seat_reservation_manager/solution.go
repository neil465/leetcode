

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
type SeatManager struct {
    h *MinHeap
}


func Constructor(n int) SeatManager {
    h := &MinHeap{}
    for i := 1 ; i <= n; i ++ {
        heap.Push(h, i)
    }
    return SeatManager{h}
}


func (this *SeatManager) Reserve() int {
    return heap.Pop(this.h).(int)
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(this.h, seatNumber)
}


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */