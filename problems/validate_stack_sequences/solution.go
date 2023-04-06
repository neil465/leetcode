func validateStackSequences(pushed []int, popped []int) bool {
    stack := []int{}
    for _, i := range pushed {
        stack = append(stack, i)
        for len(stack) > 0 && stack[len(stack) - 1] == popped[0] {
            popped = popped[1:]
            stack = stack[:len(stack) - 1]
        }
    }
    
    return len(popped) == 0
}