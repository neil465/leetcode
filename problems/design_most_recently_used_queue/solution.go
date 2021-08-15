type MRUQueue struct {
    arr []int
}


func Constructor(n int) MRUQueue {
    arr := make([]int,n)
    for i := 1 ; i <= n ; i++{
        arr[i-1] =i
    }
    return MRUQueue{arr}
}


func (this *MRUQueue) Fetch(k int) int {
    temp := this.arr[k-1]
    this.arr = append(this.arr[:k-1], this.arr[k:]...)
    this.arr = append(this.arr,temp)
    return temp
}


/**
 * Your MRUQueue object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Fetch(k);
 */