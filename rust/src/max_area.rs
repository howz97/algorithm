use std::cmp;

pub fn max_area(height: Vec<i32>) -> i32 {
    let mut max = 0;
    let mut cur_left = 0;
    let mut cur_right = height.len()-1;
    while cur_left < cur_right {
        let height_left = height[cur_left];
        let height_right = height[cur_right];
        let cur_area = (cur_right - cur_left) * (cmp::min(height_left, height_right) as usize);
        if cur_area > max {
            max = cur_area;
        }
        if height_left <= height_right {
            cur_left += 1;
        } else {
            cur_right -= 1;
        }
    }
    max as i32
}