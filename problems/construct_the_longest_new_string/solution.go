func longestString(x int, y int, z int) int {
    if y > x {

        return x * 2 + (x + 1) * 2 + z * 2
    } else if x > y {

        return (y + 1) * 2 + y * 2 + z * 2
    }
    return x * 2 + y * 2 + z *2
}