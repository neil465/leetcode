func sortArray(nums []int) []int {
    h := &IntHeap{}
    result := []int{}
    for _,i := range nums{
        heap.Push(h,i)
    }
    for h.Len() > 0{
        result = append(result,heap.Pop(h).(int))
    }
    return result
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
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }