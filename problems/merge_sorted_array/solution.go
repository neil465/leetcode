import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
    for i := 0 ; i< n ; i++{
        nums1[m+i] = nums2[i]
    }
    sort.Ints(nums1)
    
    
//     minusval := 0
    
//     if m >= n{
//         minusval = n
        
//     }else{
//         minusval = m
        
//     }
//     for i := m; i < len(nums1); i++ {
//         if i-minusval >= 0{
//             nums1[i] = nums2[i-minusval]
//         }else{
//             nums1[i] = nums2[i]
//         }
// 	}
//     sort.Ints(nums1)
}