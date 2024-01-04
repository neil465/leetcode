func selfDivisiblePermutationCount(n int) int {
    return b(make([]bool, n), n, 1)
}
func b (m []bool, n int, curInd int) int {
    s := 0
    for i := 0 ; i < n ; i++ {
        if m[i] {
            continue
        }
        if GCD(curInd, i + 1) == 1 {
            m[i] = true
            s += b(m, n, curInd + 1)
            m[i] = false
        }

    }
    if curInd == n + 1{
        s++
    }
    return s
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}