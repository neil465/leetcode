type PhoneDirectory struct {
    k map[int]bool
}


/** Initialize your data structure here
        @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
    maper := make(map[int]bool)
    for i := 0 ; i< maxNumbers ; i++{
        maper[i] = true
    }
    return PhoneDirectory{k:maper}
}


/** Provide a number which is not assigned to anyone.
        @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
    for i,j := range this.k{
        if j == true{
            this.k[i] = false
            return i
        }
    }
    return -1
}


/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
    if this.k[number] == true{
        return true
    }
    return false
}


/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int)  {
    this.k[number] = true
}


/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */