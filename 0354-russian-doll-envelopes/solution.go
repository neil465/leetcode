func maxEnvelopes(envelopes [][]int) int {
    sort.Slice(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            return envelopes[i][1] > envelopes[j][1]
        }
        return envelopes[i][0] < envelopes[j][0]
    })

    lis := []int{}

    for i := range envelopes{   
        v := sort.SearchInts(lis, envelopes[i][1])
        if v == len(lis) {
            lis = append(lis, envelopes[i][1])
        } else {
            lis[v] = envelopes[i][1]
        }

    }
    return len(lis)
}

