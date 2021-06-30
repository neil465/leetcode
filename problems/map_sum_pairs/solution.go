type MapSum struct {
    m map[string]int
}


/** Initialize your data structure here. */
func Constructor() MapSum {
    return MapSum{m : make(map[string]int)}
}


func (this *MapSum) Insert(key string, val int)  {
    this.m[key] = val
}


func (this *MapSum) Sum(prefix string) int {
    sum := 0
    for i,_ := range this.m{
        if strings.HasPrefix(i,prefix){
            sum+= this.m[i]
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