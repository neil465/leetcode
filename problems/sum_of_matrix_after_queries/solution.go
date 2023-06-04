type a struct {
    val, time int
}

func matrixSumQueries(n int, queries [][]int) int64 {
    rSum, cSum := make([]a, n), make([]a, n)
    for t, i := range queries {
        if i[0] == 0 {
            rSum[i[1]] = a{i[2], t + 1}
        } else {
            cSum[i[1]] = a{i[2], t + 1}
        }
    }

    s := 0
    for _, i := range rSum {
        for _, j := range cSum {
            if i.time > j.time {
                s += i.val
            } else {
                s += j.val
            }

        }
        
    }
    return int64(s)
}