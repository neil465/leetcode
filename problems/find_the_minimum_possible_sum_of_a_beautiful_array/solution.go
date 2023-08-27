

func minimumPossibleSum(n int, target int) int64 {
    m := map[int]bool{}
    s := 0
    c := 0
    for i := 1 ; i< 100000000; i++ {
        if c == n {
            return int64(s)
        }
        if m[target - i] {
            continue
        }
        c++
        s += i
        m[i] = true
    }
    return int64(s)
}