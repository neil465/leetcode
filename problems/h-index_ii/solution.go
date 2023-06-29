func hIndex(citations []int) int {
    
    l, r := 0, len(citations) - 1
    mid := (l + r) / 2 
    for l <= r {
        mid = (l + r) / 2 
        if citations[mid] == len(citations) - mid {

            return citations[mid]
        
        } 
        if citations[mid] < len(citations) - mid{
            l = mid + 1
        } else if citations[mid] > len(citations) - mid {
            r = mid - 1
        } 
        
    }
    if citations[mid] == len(citations) - mid {

        return citations[mid]
        
    } 
    return len(citations) - (r + 1)
}

