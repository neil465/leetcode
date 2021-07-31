import ("container/heap")
func connectSticks(sticks []int) int {
    m := &IntHeap{}
    heap.Init(m)
    sum := 0
    for _,i := range sticks{
        heap.Push(m,i)
    }
    for m.Len() > 1{

        a := heap.Pop(m)
        b := heap.Pop(m)
        sum += a.(int)+b.(int)
        heap.Push(m,a.(int)+b.(int))
   
    }
    return sum
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Pop() interface{} {
    old := *h
    pop := old[len(old)-1]
    *h = old[:len(old)-1]
    return pop
}
func (h *IntHeap) Push(x interface{}) {*h = append(*h, x.(int))}