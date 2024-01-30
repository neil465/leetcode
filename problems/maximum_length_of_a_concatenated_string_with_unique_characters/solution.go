var dp = map[[2]int]int{}

func maxLength(arr []string) int {
    dp = map[[2]int]int{}
    strBM := map[string]int{}
    startingMask := 0
    for index, str := range arr {
        mask := 0
        for _, letter := range str {
            if isSet(int(letter) - int('a'), mask) {
                startingMask = setBit(index, startingMask)
                break
            }
            mask = setBit(int(letter) - int('a'), mask)
        }
        strBM[str] = mask
    }   
    return re(strBM, startingMask, 0, arr, 0)

}
func re(strBM map[string]int, takenMask int, curStrMask int,  arr []string, curlen int) int {
    count := curlen
    if val, ok := dp[[2]int{curStrMask, takenMask}] ; ok{
        return val
    }
    for index, str := range arr {
        if isSet(index, takenMask) || curStrMask & strBM[str] > 0 {
            continue
        }
        count = max(re(strBM, setBit(index, takenMask), curStrMask | strBM[str], arr, curlen + len(str)), count)
    }
    dp[[2]int{curStrMask, takenMask}] = count
    return count
}

func setBit(k, n int) int {
    return ((1 << k) | n)
}
func unsetBit(k, n int) int {
    return (^(1 << k) & n)
}
func isSet(k, n int) bool {
    return (1 << k) & n > 0
}
