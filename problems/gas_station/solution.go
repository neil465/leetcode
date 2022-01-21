func canCompleteCircuit(gas []int, cost []int) int {
    start, amountOfGasNow, sum := 0, 0, 0
    
    for i := 0 ; i < len(gas) ; i++{
        sum += gas[i]-cost[i]
        amountOfGasNow += gas[i]-cost[i]
        
        if amountOfGasNow < 0{
            amountOfGasNow = 0
            start = i+1
        }
    }
    
    if sum < 0 {
        return -1
    }
    return start
    
}