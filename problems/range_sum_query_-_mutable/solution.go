type NumArray struct {
    a []int
}


func Constructor(nums []int) NumArray {
    return NumArray{nums}
}


func (this *NumArray) Update(index int, val int)  {
    this.a[index] = val
}


func (this *NumArray) SumRange(left int, right int) int {
    sum := 0
    for i := left ; i <= right ; i++{
        sum += this.a[i]
    }
    return sum
}



/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */