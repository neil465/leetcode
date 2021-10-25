type Trie struct {
    children [26]*Trie
    End bool
}


func Constructor() Trie {

    return Trie{End:false}
}


func (this *Trie) Insert(word string)  {
    c := this
    for _,i := range word{
        index := i-'a'
        if c.children[index] == nil{
            c.children[index] = &Trie{}
        }
        c = c.children[index]
    }
    c.End = true
}


func (this *Trie) Search(word string) bool {
    c := this
    for _,i := range word{
        index := i-'a'
        if c.children[index] == nil{
            return false
        }
        c = c.children[index]
    }
    return c.End
}


func (this *Trie) StartsWith(prefix string) bool {
    c := this
    for _,i := range prefix{
        index := i-'a'
        if c.children[index] == nil{
            return false
        }
        c = c.children[index]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
