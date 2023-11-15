type SummaryRanges struct {
    a []int
}


func Constructor() SummaryRanges {
    return SummaryRanges{[]int{}}
}


func (this *SummaryRanges) AddNum(value int)  {
    k := sort.SearchInts(this.a,  value)
    this.a = append(this.a[:k], append([]int{value}, this.a[k:]...)...)
}


func (this *SummaryRanges) GetIntervals() [][]int {
    if len(this.a) == 0 {
        return [][]int{}
    }

    k := [][]int{}
    cur := 0
    for i := 0; i < len(this.a) - 1; i++ {
        if len(k) == cur {
            k = append(k, []int{this.a[i]})
        }
        
        if this.a[i] != this.a[i + 1] - 1 && this.a[i] != this.a[i + 1] {
            k[cur] = append(k[cur], this.a[i])
            cur ++
            
        }
        
    }
    if len(k) > 0 && len(k[len(k) - 1]) == 1 {
        k[len(k) - 1] = append(k[len(k) - 1], this.a[len(this.a) - 1])
    } else{
        k = append(k, []int{this.a[len(this.a) - 1], this.a[len(this.a) - 1]})
    }
    return k
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */