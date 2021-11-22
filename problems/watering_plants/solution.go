func wateringPlants(plants []int, capacity int) int {
    waterInBucket := capacity
    steps := 0
    
    for i := 0 ; i < len(plants) ; i++{
        if plants[i] > waterInBucket{
            steps += i+i
            waterInBucket = capacity
        }
        steps++
        waterInBucket -= plants[i]
    } 
    return steps
}