func calculateTime(keyboard string, word string) int {
    m :=map[byte]int{}

    for i :=range keyboard {
        m[keyboard[i]] = i
    }

    v := m[word[0]]
    for i := 1; i< len(word); i++ {
        v += abs(m[word[i]] - m[word[i - 1]])
    }
    return v

}

func abs(i int) int {
    if i > 0 {
        return i
    }
    return - i
}