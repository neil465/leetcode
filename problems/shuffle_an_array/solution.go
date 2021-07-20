type Solution struct {
    k []int
    
}


func Constructor(nums []int) Solution {
    return Solution{nums}   
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.k
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    arr := make([]int,len(this.k))
    m := make(map[int]bool)
    rand.Seed(time.Now().UnixNano())
    
    for i:=0; i<len(this.k); i++ {
        for true{
            g := this.k[rand.Intn(len(this.k))]
            if m[g] != true{
                arr[i] = g
                m[g] = true
                break
            }
            
        }
        
    }
    
    return arr
}

