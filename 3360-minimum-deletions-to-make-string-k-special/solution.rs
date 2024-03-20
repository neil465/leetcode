use core::cmp::{max, min};
use std::collections::HashMap;
impl Solution {
    pub fn minimum_deletions(word: String, k: i32) -> i32 {
        let mut counts:HashMap<char, i32> = word.chars().fold(HashMap::new(), |mut acc, c| {
            *acc.entry(c).or_insert(0) += 1;
            acc
        });

        println!("{:?}", counts);


        let mut res:i32 = i32::MAX;
        for best_size in counts.clone().into_values() {
            let mut cur = 0;
            for val in counts.clone().into_values() {
                if val < best_size {
                    cur += val
                } else {
                    let diff_optimal_val: i32 = best_size.abs_diff(val) as i32;
                    let remove_leeway: i32 = max(diff_optimal_val - k, 0);
                    cur += remove_leeway
                }
               
            }
            res = min(res, cur)
        }
        res
    }
}


