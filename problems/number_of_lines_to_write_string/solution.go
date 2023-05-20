func numberOfLines(widths []int, s string) []int {
    v, c := 0, 0
    for _, i := range s {
        if c + widths[(i - 'a')] > 100 {
            v ++
            c = widths[(i - 'a')]
        } else {
            c += widths[(i - 'a')]
        }
    }
    return []int{v + 1, c}
}