impl Solution {
    pub fn find_peaks(mountain: Vec<i32>) -> Vec<i32> {
        mountain
        .iter()
        .enumerate()
        .filter(|&(i, &value)| {
            // Check for neighbors while avoiding out of bounds errors for edges
            i > 0 && i < mountain.len() - 1 && value > mountain[i - 1] && value > mountain[i + 1]
        })
        .map(|(i, _)| i as i32)
        .collect()
    }
}
