type mHeap struct {
	a []int
}

func connectSticks(sticks []int) int {
 
    m := &mHeap{}
    sum := 0
    for _,i := range sticks{
        m.Insert(i)
    }
    for len(m.a) > 1{

        a := m.Pop()
        b := m.Pop()
        sum += a+b
        m.Insert(a+b)
   
    }
    return   sum
}

func (h *mHeap) Pop() int{
	poped := h.a[0]
	l := len(h.a)-1
	h.a[0] = h.a[l]
	h.a = h.a[:l]
	h.HeapifyDown(0)
	return poped
}
func (h *mHeap) HeapifyDown(index int) {
    
	lastIndex := len(h.a) - 1
	l, r := Left(index), Right(index)
	Min := 0


	for l <= lastIndex {
		if l == lastIndex {
			Min = l
		} else if h.a[l] > h.a[r] {
			Min = r
		} else {
			Min = l
		}
		if h.a[index] > h.a[Min] {
			h.swap(index, Min)
			index = Min
			l, r = Left(index), Right(index)
		} else {return}

	}
}
func (h *mHeap) Insert(key int){
	h.a = append(h.a,key)
	h.HeapifyUp(len(h.a)-1)
}
func (h *mHeap) HeapifyUp(index int)  {
	for h.a[Parent(index)] > h.a[index]{
		h.swap(index,Parent(index))
		index = Parent(index)
	}
}
func Parent(index int) int{
	return (index-1)/2
}
func Left(index int) int{
	return index*2 +1
}
func Right(index int) int{
	return index*2 +2
}
func (h *mHeap) swap(i,i2 int) {
	h.a[i] , h.a[i2] = h.a[i2] , h.a[i]
}