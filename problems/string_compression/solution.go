func compress(chars []byte) int {
   var v byte
   count := 0
   res := []byte{}
   p := 0
   for _, i := range chars {
       if byte(i) != v {
           if count == 1 {
               chars[p] = v
               p ++
           } else if count > 1 {
                chars[p] = v
                p ++ 
                ag := strconv.Itoa(count)
                for _, i :=range ag {
                    chars[p] = byte(byte(i))
                    p ++
                }
               
           }
           v = byte(i)
           count = 1
       }  else {
           count++
       }
   }

     if count == 1 {
        chars[p] = v
        p ++
    } else if count > 1 {
        chars[p] = v
        p ++ 
        ag := strconv.Itoa(count)
        for _, i :=range ag {
            chars[p] = byte(byte(i))
            p ++
        }
        
    }

   fmt.Println(string(res))
   chars = res
   fmt.Println(string(chars))
   return p
}