type greatest struct{
    letter byte
    time int
}
func slowestKey(releaseTimes []int, keysPressed string) byte {
    greatarr := greatest{}
 
  
    for i := 0 ; i< len(releaseTimes) ; i++{
        fmt.Println(greatarr)
        diff := 0
        if i != 0{
            diff = releaseTimes[i]-releaseTimes[i-1]
            flag := true
            if diff == greatarr.time {
                flag = false
                if byte(keysPressed[i]) > greatarr.letter{
                    greatarr.letter = keysPressed[i]
                    greatarr.time = diff
                }
            }
            if flag {
                if diff > greatarr.time{
                    greatarr.letter = keysPressed[i]
                    greatarr.time = diff
                }
            }
            
            
        } else if releaseTimes[i] >= greatarr.time{
            greatarr.letter = keysPressed[i]
            greatarr.time = releaseTimes[i]            
        }
    }
    return greatarr.letter
    
}