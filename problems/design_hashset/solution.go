type MyHashSet struct {
    k   map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    x:= MyHashSet{}
    x.k = make(map[int]int)
    return x
}


func (this *MyHashSet) Add(key int)  {
    this.k[key] = 1
}


func (this *MyHashSet) Remove(key int)  {
    this.k[key] = 0
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    if this.k[key] == 1{
        return true
    }

    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */