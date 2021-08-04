func maximumPopulation(logs [][]int) int {

    a := make(map[int]int)
    res := &IntHeap{}
    
    for _,i := range logs{
        for j := i[0] ;  j < i[1] ; j++{
            a[j]++
            heap.Push(res,j)
        }
    }
    
    max := 0
    date := 0
   
    
    for res.Len() > 0{
        poped := heap.Pop(res).(int)
        if max < a[poped]{
            
            max = a[poped]
            date = poped
        }
    }
    return date
    
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