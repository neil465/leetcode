func removeDuplicates(S string) string {
    for i := 0 ; i < len(S)-1 ; i++{
        
        if S[i] == S[i+1]{
            S = S[:i] + S[i+2:]
            i = -1 
        }
    }
    
    return S
}