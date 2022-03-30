func searchMatrix(matrix [][]int, target int) bool {
    i,j := 0 , 0 
    oldi,oldj := -1,-1
    for i != oldi || j != oldj{
        oldi , oldj = i,j
        if i != len(matrix)-1 && target >= matrix[i+1][j]{
            i++
        } else if j != len(matrix[0])-1 && target >= matrix[i][j+1]{
            j++
        }
        if matrix[i][j] == target{
            return true
        }
        
        
    }
    if matrix[i][j] == target{
        return true
    }
    return false
}