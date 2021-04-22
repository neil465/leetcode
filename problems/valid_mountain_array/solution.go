import "fmt"
func validMountainArray(arr []int) bool {
    flag := true
    if len(arr) == 1{
        return false
    }
	for i := 0; i < len(arr)-1; i++ {
		
		if arr[i] == arr[i+1] {
			return false
		}
		if arr[i] > arr[i+1]{
			if i==0 {
				return false
			}else {
                fmt.Println("down")
				flag = false
                
			}
		}
        fmt.Println(flag)
		if arr[i] < arr[i+1] && flag == false{
            
            fmt.Println("up")
            return false
        }
      
	}
    if len(arr)>1 && arr[len(arr)-1]>= arr[len(arr)-2]  {
		return false
	}
	return true
}
