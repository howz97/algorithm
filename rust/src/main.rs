use std::collections::HashMap;

// fn num_way(n: i32) -> i32 {
//     if n == 0 || n == 1 {
//         return 1
//     }
//     let mut num0 = 1;
//     let mut num1 = 1;
//     let mut tmp;
//     let mut idx_num1 = 1;
//     while idx_num1 < n {
//         idx_num1 += 1;
//         tmp = (num1 + num0) % 1000000007;
//         num0 = num1;
//         num1 = tmp;
//     }
//     return num1
// }

// fn min_array(numbers: Vec<i32>) -> i32 {
//     let first = numbers[0];
//     let mut last = first;
//     for e in numbers.iter() {
//         let v = *e;
//         if v < last {
//             return v
//         } else {
//             last = v
//         }
//     }
//     return first
// }

// fn length_of_longest_substring(s: String) -> i32 {
//     let mut longest = 0;
//     let mut start = 0;
//     let mut end = 0;
//     while end < s.len() {
//         end += 1;
//         let idx_r = idx_repeat(s[start..end].to_string());
//         if idx_r < 0 {
//             longest = end - start;
//         } else {
//             let uidx = idx_r as usize;
//             start += uidx + 1;
//             end += uidx;
//         }
//     }
//     return longest as i32
// }

// fn idx_repeat(sub: String) -> i32 {
//     let mut h = HashMap::new();
//     let mut idx = sub.len() - 1;
//     for e in sub.chars().rev() {
//         if h.contains_key(&e) {
//             return idx as i32
//         }
//         h.insert(e, true);
//         idx -= 1;
//     }
//     return -1
// }

fn main() {
    let mut numbers = vec![0,1,1,2,3,4,5,5,5,6,7,8,9,0,0];
    let mean = calc_mean(&numbers);
    match mean {
        Some(i) => println!("mean is {}", i),
        None => println!("mean not exist")
    }
    let median = calc_median(&mut numbers);
    match median {
        Some(i) => println!("median is {}", i),
        None => println!("median not exist")
    }
    let mode = calc_mode(&numbers);
    match mode {
        Some(i) => println!("mode is {}", i),
        None => println!("mode not exist")
    }
}

fn calc_mean(numbers: &Vec<i32>) -> Option<i32> {
    if numbers.len() == 0 {
        return None
    }
    let mut sum = 0;
    for number in numbers {
        sum += number
    }
    Some(sum/(numbers.len() as i32))
}

fn calc_median(numbers: &mut Vec<i32>) -> Option<&i32> {
    let length = numbers.len();
    if length % 2 == 0 {
        return None
    }
    numbers.sort();
    let median_idx = length / 2;
    numbers.get(median_idx)
}

fn calc_mode(numbers: &Vec<i32>) -> Option<i32> {
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