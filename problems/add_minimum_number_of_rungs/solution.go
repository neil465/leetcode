func addRungs(rungs []int, dist int) int {
    rungs = append([]int{0}, rungs...)
    val := 0
    for i := 0 ; i < len(rungs) - 1; i++ {
        val += ((rungs[i + 1] - rungs[i]) / dist) 
        if (rungs[i + 1] - rungs[i]) % dist == 0 {
            val --
        }
    }
    return val
}