func hasAlternatingBits(n int) bool {
	state := 0
	if (n & 1 == 0)  {	
		state = 1
	}
	for n>0 {
		if (n & 1) == state  {
			return false
		}
        if (state == 0){
            state = 1
        }else{
            state = 0
        }
		n>>=1
	}
	return true
	
	
}
