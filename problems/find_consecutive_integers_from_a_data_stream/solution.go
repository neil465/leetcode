type DataStream struct {
    s int
    v int
    k int
}


func Constructor(value int, k int) DataStream {
    return DataStream{0, value, k}
}


func (this *DataStream) Consec(num int) bool {


    if num != this.v {

        this.s = 0
    } else {
        this.s ++
    }

    


    return this.s >= this.k

}


/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */