func isOneBitCharacter(bits []int) bool {
    if len(bits) <= 2 {
        return bits[0] != 1
    }
    if bits[0] == 1 {
        return isOneBitCharacter(bits[2:])
    }
    return isOneBitCharacter(bits[1:])
}