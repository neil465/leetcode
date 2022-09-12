func mostFrequentEven(nums []int) int {
    var m [100000]int
    max, num  := 0, -1     
    for _, i := range nums {
        m[i] ++
        if i % 2 == 0 {
            if m[i] > max {
                max = m[i]
                num = i
            } else if m[i] == max && i < num {
                num = i
            }
        } 
    }
    return num
}