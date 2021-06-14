type MovingAverage struct {
    k []int
    j int
    
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{k: []int{},j:size}
}



func (this *MovingAverage) Next(val int) float64 {
    if len(this.k) == this.j{
        this.k = this.k[1:]
    }
    this.k = append(this.k,val)
    sum :=0
    for _,i := range this.k{
        sum += i
    }
    
    return float64(sum)/float64(len(this.k))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */