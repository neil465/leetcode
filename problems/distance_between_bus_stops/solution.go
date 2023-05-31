func distanceBetweenBusStops(distance []int, start int, destination int) int {
    runningSum := 0
    t := 0
    for i := start; i < len(distance) + start; i ++ {
        if i % len(distance) == destination {
            t = runningSum
        }
        runningSum += distance[i % len(distance)]
    }
    return int(math.Min(float64(t), float64(runningSum - t)))
}