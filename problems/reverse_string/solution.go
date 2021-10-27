func reverseString(s []byte)  {
    helper(0,len(s)-1,&s) 
}
func helper (i int ,j int ,s *[]byte){
	if i >= j{return}
	(*s)[i], (*s)[j] = (*s)[j],(*s)[i]
	helper(i+1,j-1,s)
}