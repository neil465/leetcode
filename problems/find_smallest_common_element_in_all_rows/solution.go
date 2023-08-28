func smallestCommonElement(mat [][]int) int {
    k := [10001]int{}
    for _, i := range mat {
      for _, j := range i {
        k[j] ++
        if k[j] == len(mat) {
          return j
        }
      }
    }

    return -1
}