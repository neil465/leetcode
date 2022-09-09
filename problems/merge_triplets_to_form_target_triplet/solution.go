func mergeTriplets(triplets [][]int, target []int) bool {
    var temp [3]int
    for _, i := range triplets {
        if i[0] <= target[0] && i[1] <= target[1] && i[2] <= target[2] {
            temp[0], temp[1], temp[2] = max(i[0],temp[0]) , max(i[1],temp[1]), max(i[2], temp[2])
        } 
    }
    return temp[0] == target[0] && temp[1] == target[1] && temp[2] == target[2]
}

func max (i, j int) int { 
    if i > j {
        return i 
    } else { 
        return j 
    } 
}