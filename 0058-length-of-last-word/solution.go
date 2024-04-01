func lengthOfLastWord(s string) int {
    end := len(s) - 1

    for s[end] == ' ' && end > 0 {
        end --
    }
    v := 0

    for s[end] != ' ' && end > 0{
        end --
        v ++
    }

    if s[end] != ' ' {
        v ++
    }

    return v

    
}
