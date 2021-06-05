type ProductOfNumbers struct {
    arr []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{}
}


func (this *ProductOfNumbers) Add(num int)  {
    this.arr = append(this.arr,num)
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    res := 1
    count := len(this.arr)-1
    for i:= 0 ; i<k ; i++{
        res *= this.arr[count]
        count--
    }
    return res
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */