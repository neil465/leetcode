func countBeautifulPairs(nums []int) int {
    b := 0
    for i := range nums {
        for j := i + 1; j < len(nums); j++{
            c, d := nums[j], nums[i]
            ab, bc := c % 10, 0

            for d > 0 {
                bc = d % 10
                d /= 10
            }
            if hcf(ab, bc) == 1 {
                b++
            }
        }
    }
    return b
}
func hcf(n1 int, n2 int) int {
   if (n2 != 0) {
      return hcf(n2, n1 % n2);
   } else {
      return n1;
   }
}