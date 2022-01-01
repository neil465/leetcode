var maxsum = -100000
func maxSubArray(nums []int) int {
    maxsum = -100000
    var calculate func([]int, int) int
    calculate = func(nums []int,pos int) int {
        if pos == 0 {
            maxsum = int(math.Max(float64(maxsum),float64(nums[0])))
            return nums[0]
        }
        minisum := int(math.Max(float64(nums[pos]),float64(nums[pos]+calculate(nums,pos-1))))
        maxsum = int(math.Max(float64(maxsum),float64(minisum)))
        return minisum
    }
    calculate(nums,len(nums)-1)

    return maxsum
}