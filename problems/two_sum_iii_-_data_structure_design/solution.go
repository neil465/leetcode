type TwoSum struct {
    k []int
}


/** Initialize your data structure here. */
func Constructor() TwoSum {
    return TwoSum{}
}


/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int)  {
    this.k = append(this.k,number)
}


/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
    sort.Ints(this.k)
    fmt.Println(this.k,value)
    l := 0
    r := len(this.k)-1
    if len(this.k) == 0 ||len(this.k) == 1{
        return false
    }
    for l <r{
        if this.k[l] + this.k[r] == value{
            return true
        }
        if  this.k[l] + this.k[r] > value{
            r--
        }
        if this.k[l] + this.k[r] < value{
            l++
        }
    }
    return false
}


/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */