func binaryGap(n int) int {
    sum, temp := 0, 0
    a := strconv.FormatInt(int64(n), 2)

    a = strings.Trim(a, "0")

    if strings.Count(a, "1") <2 {
        return 0
    }

    for i := 0 ; i < len(a); i++ {
        if a[i] == '0' {
            temp ++
        } else {
            if temp > sum {
                sum = temp
            }
            temp = 0
        }
    }
    return sum + 1

}