func passThePillow(n int, time int) int {
    time = time % (n * 2 - 2)

    if time >= n {
        return n - (time - n + 1)
    } 
    return time + 1
}