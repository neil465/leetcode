func brokenCalc(startValue int, target int) int {
    var result int
    
    for target > startValue {
        
        // we are adding to result which is the number of steps we have taken, and we are adding to target, and 
        // assuming that target is not divisible by 2  
        result, target = result+1, target + 1
        
        if ( target - 1 ) % 2 == 0 { // if target is divisible by 2 we remove the one we added to target and divide 
            //target by 2
            target = ( target - 1 ) / 2
        }
    }
    
    return result + startValue - target
}