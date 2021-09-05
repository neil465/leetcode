type MapSum struct {
    m map[string]int
}


/** Initialize your data structure here. */
func Constructor() MapSum {
    a := make(map[string]int)
    return MapSum{a}
}


func (this *MapSum) Insert(key string, val int)  {
    this.m[key] = val
}


func (this *MapSum) Sum(prefix string) int {
    sum:= 0
    for j,i := range this.m{
        if strings.HasPrefix(j, prefix){
            sum+= i
        } 
    }
    return sum
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */