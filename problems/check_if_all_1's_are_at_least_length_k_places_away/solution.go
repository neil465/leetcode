func kLengthApart(nums []int, k int) bool {
    flag , count := false,0
    for i := 0 ; i < len(nums); i ++{
        if nums[i] == 1 && flag == false{
            fmt.Println("1",count)
            flag = true
            count = 0 
            continue
        }else if nums[i] == 1 && flag == true{
            fmt.Println("2",count)
            if count >= k{
                count = 0
                continue
            }else{
                return false   
            }
            
        }
        count++
    }
    return true
}