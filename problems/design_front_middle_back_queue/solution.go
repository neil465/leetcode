type FrontMiddleBackQueue struct {
    a []int
}


func Constructor() FrontMiddleBackQueue {
    return FrontMiddleBackQueue{[]int{}}
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {

    this.a = append([]int{val},this.a...)

}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {

    this.a = append(this.a[:len(this.a)/2 ], append([]int{val}, this.a[len(this.a)/2:]...)...)


}


func (this *FrontMiddleBackQueue) PushBack(val int)  {

    this.a = append(this.a, val)


}


func (this *FrontMiddleBackQueue) PopFront() int {

    if len(this.a) == 0 {
        return -1
    }
    v := this.a[0]
    this.a = this.a[1:]


    return v
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
    

    if len(this.a) == 0 {
        return -1
    }
    if len(this.a) % 2 == 0 {
        v := this.a[len(this.a) / 2 -1]
        this.a = append(this.a[:len(this.a)/2 - 1], this.a[len(this.a)/2:]...)
        return v
    }
    v := this.a[len(this.a) / 2]
    this.a = append(this.a[:len(this.a)/2], this.a[len(this.a)/2 + 1:]...)


    return v
}


func (this *FrontMiddleBackQueue) PopBack() int {

    if len(this.a) == 0 {
        return -1
    }
    v := this.a[len(this.a) - 1]
    this.a = this.a[:len(this.a) - 1]


    return v
}


/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */