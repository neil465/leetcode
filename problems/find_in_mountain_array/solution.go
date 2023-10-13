/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, ma *MountainArray) int {
    start := 0
    if ma.get(start) == target {
        return 0
    }
    l := ma.length() -1 
    end := l
    
    mid := 0
    for start < end {
        mid = (end + start)/2 

        if mid == 0 || mid == l {
            
            break
        }
        bef, cur, aft := ma.get(mid - 1), ma.get(mid), ma.get(mid + 1)
        
        if bef < cur && aft < cur {
            
            break
        }   
        if bef > cur && cur > aft {
            end = mid 
        } else {
            start = mid +1
        }
    }

    if ma.get(mid) == target {
        return mid
    }
    mountainMid := mid
    start = 0
    end = max(mountainMid , 0 )
    mid = 0
    for start < end {
        mid = (end + start)/2
        cur := ma.get(mid)
        
        if cur == target {
            break
        }   
        if cur > target {
            end = mid 
        } else {
            start = mid +1
        }
    }
    if ma.get(mid) == target {
        return mid
    }

    start = min(mountainMid , ma.length() -1 )
    end = ma.length() -1 
    mid = 0
    for start < end {
        mid = (end + start)/2
        cur := ma.get(mid)
        if cur == target {
            break
        }   
        if cur < target {
            end = mid
        } else {
            start = mid +1
        }
    }
    if ma.get(mid) == target {
        return mid
    }
    if ma.get(l) == target {
        return l
    }

    return -1
}
func min(i,j int)int {
    if i < j {
        return i
    }
    return j
}
func max(i,j int)int {
    if i > j {
        return i
    }
    return j
}