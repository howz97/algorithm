// use std::collections::HashMap;

fn main() {
    // let ve = vec![2,2,2,0,1];
}

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
