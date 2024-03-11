impl Solution {
    pub fn minimum_boxes(apple: Vec<i32>, capacity: Vec<i32>) -> i32 {
        let mut sum = 0;

        for v in apple.iter() {
            sum += v;
        }
        let mut c =capacity.clone();
        c.sort_by_key(|n| std::i32::MAX - n);


        for (ind, v) in c.iter().enumerate() {
            if sum <= 0 {
                return ind as i32;
            }
            sum -= v;
        }
        return c.len() as i32
    }
}
