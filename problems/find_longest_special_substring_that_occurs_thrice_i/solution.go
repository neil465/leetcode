func maximumLength(s string) int {
    i, j := 0, 0

    cur := string(s[0])
    m := map[string]int{}
    for j < len(s) - 1  {

        if s[j + 1] != cur[0] {

            for iter := 1 ; iter <= j - i + 1; iter ++ {
                m[cur[:len(cur) - iter + 1]] += iter
            }
            i = j + 1
            j ++
            cur = string(s[j ])
        } else {
            j ++
            cur += string(s[j])
        }
    }
    for iter := 1 ; iter <= j - i + 1; iter ++ {
        m[cur[:len(cur) - iter + 1]] += iter
    }

    ma := -1

    for k, v := range m {
        if v >= 3 {
            ma = max(len(k), ma)
        }
    }
    return ma
}