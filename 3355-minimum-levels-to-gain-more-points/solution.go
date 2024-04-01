func minimumLevels(possible []int) int {
    psum := []int{}

    for _, i := range possible {
        if len(psum) == 0 {
            if i > 0 {
                psum = append(psum, 1)
            } else {
                psum = append(psum, -1)
            }
            
        } else {
            if i > 0 {
                psum = append(psum, 1 + psum[len(psum) -1])
            } else {
                psum = append(psum, -1 + psum[len(psum) -1])
            }
        }
    }
    for i := 0 ; i < len(psum) - 1; i ++ {
        if psum[i] > psum[len(psum) - 1] - psum[i] {
            return i + 1
        }
    }
    return -1
}
