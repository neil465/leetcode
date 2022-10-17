package main

import "regexp"
// In this solution, we use backtracking to find all the possible times. We go through 
// all the '?' characters and replace them with numbers from 0 to 9. We check if any of the times are valid,
// if we find a valid time, we increment our sum. We do this for all the '?' characters.

// walkthrough of the solution given the following input: "2?:00"
// time = "2?:00", i = 0
// time = "2?:00", i = 1
// we have now found a '?', so we replace it with all the numbers from 0 to 9
// time = "20:00", i = 2 // valid time found
// time = "21:00", i = 2 // valid time found
// time = "22:00", i = 2 // valid time found
// time = "23:00", i = 2 // valid time found
// time = "24:00", i = 2 // invalid time found
// time = "25:00", i = 2 // invalid time found
// time = "26:00", i = 2 // invalid time found
// time = "27:00", i = 2 // invalid time found
// time = "28:00", i = 2 // invalid time found
// time = "29:00", i = 2 // invalid time found
// time = "2?:00", i = 2
// time = "2?:00", i = 3
// time = "2?:00", i = 4

// we return 3, for we have found three valid times
func countTime(time string) int {
   return t([]rune(time), 0)
}
func t(time []rune, i int) int { // helper function to calculate the number of valid times
   // we check if the time is valid
   re := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$`)
   if re.MatchString(string(time)) {
      return 1
   }
   // we check if we are at the end of the time, and we have not found a valid time
   if i >= 5 {
      return 0
   }

   a := 0
   if time[i] == '?' { // if the current character is a '?'
      for _, j := range "0123456789" { // we iterate through all the possible numbers
         // backtracking with the current number and replacing the '?' with numbers from 0 to 9
         time[i] = j
         a += t(time, i+1)
         time[i] = '?'
      }

   } else { // if the current character is not a '?'
      // we continue with the next character
      a = t(time, i+1)
   }
   // we return the number of valid times we have found if any
   return a
}
