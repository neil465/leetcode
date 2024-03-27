impl Solution {
    pub fn count_tested_devices(battery_percentages: Vec<i32>) -> i32 {
        let mut cnt = 0;
        for v in battery_percentages.iter() { cnt += if v - cnt > 0 {1} else {0}}
        cnt
    }
}
