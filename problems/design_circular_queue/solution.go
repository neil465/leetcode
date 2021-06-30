type MyCircularQueue struct {
    l int
    a []int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{l:k , a:[]int{}}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.l == len(this.a){
        return false
    }
    this.a = append(this.a,value)
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if len(this.a) == 0{
        return false
    }
    this.a = this.a[1:]
    return true
}


func (this *MyCircularQueue) Front() int {
    if len(this.a) > 0{
        return this.a[0]
    }
    return -1
}


func (this *MyCircularQueue) Rear() int {
    if len(this.a) > 0{
        return this.a[len(this.a)-1]
    }
    return -1
}


func (this *MyCircularQueue) IsEmpty() bool {
    if len(this.a) > 0{
        return false
    }
    return true
}


func (this *MyCircularQueue) IsFull() bool {
    if len(this.a) == this.l{
        return true
    }
    return false
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */