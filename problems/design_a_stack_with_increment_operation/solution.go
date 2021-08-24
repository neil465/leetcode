type CustomStack struct {
    a []int
    max int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{max:maxSize}
}


func (this *CustomStack) Push(x int)  {
    if len(this.a) < this.max{this.a = append(this.a,x)}
}


func (this *CustomStack) Pop() int {
    if len(this.a) == 0{return -1}
    a := this.a[len(this.a)-1]
    this.a = this.a[:len(this.a)-1]
    return a
}


func (this *CustomStack) Increment(k int, val int)  {
    if len(this.a) <= k{
        for i,_ := range this.a{
            this.a[i] = this.a[i]+val
        }
    }else{
        for i := 0 ; i < k ; i++{
            this.a[i] = this.a[i]+val
        }  
    } 
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */