func minProcessingTime(processorTime []int, tasks []int) int {
    sort.Ints(tasks)
    sort.Sort(sort.Reverse(sort.IntSlice(processorTime)))

    s := 0
    pi := 0
    for i := 0 ; i < len(tasks); i, pi = i + 4, pi + 1 {
        m := tasks[i]
        for j := i + 1; j < i + 4; j ++ {
            if m < tasks[j] {
                m = tasks[j]
            }
        }

        if s < processorTime[pi] + m{
            s = processorTime[pi] + m
        }

    }
    return s
}