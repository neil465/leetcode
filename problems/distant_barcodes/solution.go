type pair struct {
    bar, val int
}

type Heap []pair
func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rearrangeBarcodes(barcodes []int) []int {
    m, h, res := map[int]int{}, &Heap{}, []int{}
    heap.Init(h)
    for _, i := range barcodes { m[i] ++ }
    for k, v := range m { heap.Push(h, pair{k, v}) }
    
    for h.Len() > 0 {
        left := []pair{}
        
        for {
            if h.Len() == 0 {
                break
            }
            v := heap.Pop(h).(pair)
            left = append(left, v)
            if len(res) == 0 || v.bar != res[len(res) -1 ] {
                break
            }
        }
        for i := 0 ; i < len(left); i++ {
            if i == len(left) - 1 {
                left[i].val --
                res = append(res, left[i].bar)
            } 
            if left[i].val == 0 {
                continue
            }
            heap.Push(h, left[i])
        }

    }
    return res
}