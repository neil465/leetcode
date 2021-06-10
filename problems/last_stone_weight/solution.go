func lastStoneWeight(stones []int) int {
    
    for len(stones) >1 {
        sort.Ints(stones)
        if stones[0] == stones[1] && len(stones) == 2{
            return 0
        }
        if stones[len(stones)-1] == stones[len(stones)-2]{
            fmt.Println(stones)
            stones = stones[:len(stones)-1]
            fmt.Println(stones)
            stones = stones[:len(stones)-1]
            fmt.Println(stones)
            
        }else{
            stones[len(stones)-1] = stones[len(stones)-1] - stones[len(stones)-2]
            stones = append(stones[:len(stones)-2], stones[len(stones)-1:]...)
        }
        
        
    }
    return stones[0]
}