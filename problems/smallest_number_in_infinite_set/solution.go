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


type SmallestInfiniteSet struct {
    values map[int]int
    h *IntHeap
}


func Constructor() SmallestInfiniteSet {
    h := &IntHeap{}
    heap.Init(h)
    a := map[int]int{}
    for i := 1 ; i < 1001; i ++ {
        heap.Push(h, i)
        a[i] ++
    }   
    return SmallestInfiniteSet{a, h}
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    if this.h.Len() == 0 {
        return  -1
    }
    p := heap.Pop(this.h).(int)
    this.values[p]--
    return p
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if this.values[num] == 0{
        heap.Push(this.h, num)
        this.values[num] ++
    }
    
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */