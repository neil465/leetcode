var a [1000]int
func combinationSum4(nums []int, target int) int {
    a = [1000]int{}
    for i:=0;i<1000;i++{
        a[i] = -1
    }
    return c(nums,target)
}
func c(nums []int, target int) int {
    if a[target] != -1 {
        return a[target]
    }
    if target == 0 {
        return 1
    }
    s := 0
    for _,i:= range nums {
        if target >= i {
            s += c(nums, target - i)
        }
    }
    a[target]= s
    return a[target]
}