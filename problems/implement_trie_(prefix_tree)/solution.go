type Node struct{
    children [26]*Node
    isEnd bool
}
type Trie struct {
    root *Node
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{root:&Node{}}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    currNode := this.root
    for i := 0 ; i < len(word) ; i++{
		charIndex := word[i] - 'a'
        
		if currNode.children[charIndex] == nil{currNode.children[charIndex] = &Node{}}
		currNode = currNode.children[charIndex]
	}
    currNode.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    currNode := this.root
    
	for i := 0 ; i < len(word) ; i ++{
		charIndex := word[i]-'a'
		if currNode.children[charIndex] == nil{return false}
		currNode = currNode.children[charIndex]
	}
    
    if currNode.isEnd == true{return true}
    return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    currNode := this.root
    count := 0
	for i := 0 ; i < len(prefix) ; i ++{
		count++
        charIndex := prefix[i]-'a'
		if currNode.children[charIndex] == nil{return false}
		currNode = currNode.children[charIndex]
        
        if count == len(prefix){return true}
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