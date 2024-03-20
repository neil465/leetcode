impl Solution {
    pub fn count_substrings(s: String, c: char) -> i64 {
         let v = s.chars().filter(|&t| t == c).count() as i64;
        (v * (v + 1))/2
    }
}
