func reversePrefix(word string, ch byte) string {
    a := strings.IndexByte(word,ch)
    if a == -1{return word}
    l,r := 0,a
    words := strings.Split(word,"")
    for l <= r{
        words[l],words[r] = words[r],words[l]
        l++
        r--
    }
    return strings.Join(words,"")
}