impl Solution {
    pub fn is_substring_present(s: String) -> bool {
        let rev_s: String = s.chars().rev().collect();
        let set_s: std::collections::HashSet<_> = s.chars().zip(s.chars().skip(1)).collect();
        let set_rev_s: std::collections::HashSet<_> = rev_s.chars().zip(rev_s.chars().skip(1)).collect();
        !set_s.is_disjoint(&set_rev_s)
    }
}

