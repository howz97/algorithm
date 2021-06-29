use std::collections::HashMap;

pub fn calc_mean(numbers: &Vec<i32>) -> Option<i32> {
    if numbers.len() == 0 {
        return None
    }
    let mut sum = 0;
    for number in numbers {
        sum += number
    }
    Some(sum/(numbers.len() as i32))
}

pub fn calc_median(numbers: &mut Vec<i32>) -> Option<&i32> {
    let length = numbers.len();
    if length % 2 == 0 {
        return None
    }
    numbers.sort();
    let median_idx = length / 2;
    numbers.get(median_idx)
}

pub fn calc_mode(numbers: &Vec<i32>) -> Option<i32> {
    if numbers.len() == 0 {
        return None
    }
    let mut count_map = HashMap::new();
    for number in numbers {
        let count = count_map.entry(*number).or_insert(0);
        *count += 1;
    }
    let mut count_max = 0;
    let mut mode = 0;
    for (number, count) in count_map {
        if count > count_max {
            count_max = count;
            mode = number;
        }
    }
    Some(mode)
}



