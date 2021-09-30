type OrderedStream struct {
    p int
    a []string
}


func Constructor(n int) OrderedStream {
    return OrderedStream{0,make([]string,n)}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
    a := []string{}
    this.a[idKey-1] = value
    for this.p < len(this.a) &&this.a[this.p] != ""{
        a = append(a,this.a[this.p])
        this.p++
    }
    return a
    
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */