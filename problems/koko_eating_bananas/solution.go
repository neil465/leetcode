func minEatingSpeed(piles []int, h int) int {
    
    left,right,mid := 0,0,0
    
    for _,i := range piles{
        if i > right {
            right = i
        }   
    }
    
    for left < right{
        if left+1 == right{
            return right
        }
        mid = left + (right - left) / 2
        timePast := 0
        for _,i := range piles{
            timePast += i/mid
            if i % mid != 0{
                timePast++
            }
        }
        if timePast > h {
            left = mid
        }else{
            right = mid
        }

    }
    return mid
}