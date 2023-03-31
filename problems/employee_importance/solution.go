/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    m := map[int]*Employee{}

    for _, i := range employees {
        m[i.Id] = i
    }    
    a := find(m, id) + m[id].Importance

    return a
}

func find(m map[int]*Employee, id int) int {
    sum := 0 
    for _, i := range m[id].Subordinates {
        sum += m[i].Importance + find(m, i)

    }
    return sum
}