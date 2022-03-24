func minimumHealth(damage []int, armor int) int64 {
    healthNeeded := 0 
    maxDamage := 0 
    for _,i := range damage {
        healthNeeded += i
        if i > maxDamage {
            maxDamage = i
        }
    }
    if armor > maxDamage {
        armor = maxDamage
    }
    return int64(healthNeeded - armor + 1) 
}