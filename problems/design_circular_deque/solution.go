type MyCircularDeque struct {
    maxLen int
    a []int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{k, []int{}}
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    
    if len(this.a) < this.maxLen {
        this.a = append([]int{value}, this.a...)
        

        return true
    }
    return false
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if len(this.a) < this.maxLen {
        this.a = append(this.a, value)

        return true
    }
    return false
}


func (this *MyCircularDeque) DeleteFront() bool {
    if len(this.a) > 0 {
        this.a = this.a[1:]

        return true
    }
    return false
}


func (this *MyCircularDeque) DeleteLast() bool {
    if len(this.a) > 0 {
        this.a = this.a[:len(this.a) - 1]

        return true
    }
    return false
}


func (this *MyCircularDeque) GetFront() int {
    if len(this.a) > 0 {

        return this.a[0]
        
    }
    return -1
}


func (this *MyCircularDeque) GetRear() int {
    if len(this.a) > 0 {

        return this.a[len(this.a) - 1]
    }
    return -1
}


func (this *MyCircularDeque) IsEmpty() bool {

    return len(this.a) == 0
}


func (this *MyCircularDeque) IsFull() bool {

    return len(this.a) == this.maxLen
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */