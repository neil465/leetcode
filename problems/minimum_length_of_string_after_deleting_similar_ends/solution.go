func minimumLength(s string) int {
    i, j := 0, len(s) - 1 

    for i < j{

        if s[i] != s[j] {
            return j - i + 1
        }
        for i < j {
            if s[i] == s[i + 1] {
                i ++
            } else {
                break
            }
        }
        for i < j {
            if s[j] == s[j - 1] {
                j --
            } else {
                break
            }
        }
        
        i ++
        j --
    }
    if i > j {
        return 0
    }
    return j - i + 1
}