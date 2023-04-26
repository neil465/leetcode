func shiftingLetters(s string, shifts []int) string {
    add := 0
    aa := []byte(s)
    for _, i := range shifts {
        add += i
    }

    for i := range s {
        aa[i] = rotate(aa[i], add)
        if add > 0 {
            add -= shifts[i]
        }
    }
    return string(aa)
}

func rotate(a byte, push int) byte {
    return byte(((int(a) - 97 + push) % 26) + 97)
}