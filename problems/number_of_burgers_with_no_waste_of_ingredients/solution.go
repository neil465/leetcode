func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
   
    j := (tomatoSlices - 2 *cheeseSlices ) / 2 
    s := cheeseSlices - j
    if (j + s != cheeseSlices || 4*j + 2 * s != tomatoSlices ) || j < 0 || s < 0{
        return []int{}
    }
    return []int{j, s}
}