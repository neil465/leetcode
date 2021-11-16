func timeRequiredToBuy(tickets []int, k int) int {
    count := 0
    for {
        for i := 0 ; i < len(tickets) ; i++{
            if tickets[i] == 0{continue}
            tickets[i]--
            count++
            if tickets[k] == 0{return count}
        }
        
    }   
}