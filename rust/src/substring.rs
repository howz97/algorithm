use std::collections::HashMap;

pub fn length_of_longest_substring(s: String) -> i32 {
    let mut longest = 0;
    let mut start = 0;
    let mut end = 0;
    while end < s.len() {
        end += 1;
        let idx_r = idx_repeat(s[start..end].to_string());
        if idx_r < 0 {
            longest = end - start;
        } else {
            let uidx = idx_r as usize;
            start += uidx + 1;
            end += uidx;
        }
    }
    return longest as i32
}

fn idx_repeat(sub: String) -> i32 {
    let mut h = HashMap::new();
    let mut idx = sub.len() - 1;
    for e in sub.chars().rev() {
        if h.contains_key(&e) {
            return idx as i32
        }
        h.insert(e, true);
        idx -= 1;
    }
    return -1
}