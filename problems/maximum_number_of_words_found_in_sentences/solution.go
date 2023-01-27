func mostWordsFound(sentences []string) int {
    max := 0

    for _, i := range sentences {
        temp := 1

        for _, r := range i {
            if string(r) == " " {
                temp++
            }
        }

        max = int(math.Max(float64(temp), float64(max)))

    }

    return max
}