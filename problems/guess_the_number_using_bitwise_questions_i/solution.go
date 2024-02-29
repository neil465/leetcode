/** 
 * Definition of commonSetBits API.
 * func commonSetBits(num int) int;
 */

func findNumber() int {
    v := 0
    for i := 0 ; i < 31; i ++ {
        if commonSetBits(1 << i) == 1 {
            v |= 1 << i
        }
    }
    return v
}