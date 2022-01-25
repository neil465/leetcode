func validMountainArray(arr []int) bool {
    if len(arr) < 3 { return false }
    
    left := 0
    right := len(arr)-1 
    
    for left = 0 ; left+1 <len(arr) && arr[left] < arr[left+1];left++ {
        continue 
    }
    for right = len(arr)-1 ; right > 0 && arr[right] < arr[right-1]; right--{
        continue 
    }
    return left == right && left != 0 && right != len(arr) - 1
}