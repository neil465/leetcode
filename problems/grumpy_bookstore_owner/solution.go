func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    s := 0
    for i := range customers {
        if grumpy[i] == 0 {
            s += customers[i]
        }
    }
    m := s

    for i := 0 ; i < len(grumpy) - minutes + 1; i ++ {
        t := s
        for j := i ; j < i + minutes; j ++ {
            if grumpy[j] == 1 {
                t += customers[j]
            }
        }
        if t > m {
            m = t

        }
    }

    return m
}