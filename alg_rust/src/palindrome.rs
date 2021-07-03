use std::str::FromStr;

pub fn longest_palindrome(s: String) -> String {
    if s.len() == 1 {
        return s
    }
    let str1 = search_longest_palindrome(&s, false);
    let str2 = search_longest_palindrome(&s, true);
    let result;
    if str1.len() > str2.len() {
        result = match String::from_str(str1) {
            Ok(r) => r,
            Err(e) => panic!("{}", e)
        };
    } else {
        result = match String::from_str(str2) {
            Ok(r) => r,
            Err(e) => panic!("{}", e)
        };
    }
    result
}

fn search_longest_palindrome(s: &str, start_single: bool) -> &str {
    let mut left = 0;
    let mut right = 1;
    if !start_single {
        right = 2;
    }
    let mut left_border = 0;
    let mut right_border = 1;
    loop {
        if is_palindrome(s[left..right].as_bytes()) {
            left_border = left;
            right_border = right;
            if left == 0 {
                left += 1;
                right += 1;
            }
            left -= 1;
            right += 1;
            if right > s.len() {
                break
            }
        } else {
            left += 1;
            right += 1;
            if right > s.len() {
                break
            }
        }
    }
    &s[left_border..right_border]
}

fn is_palindrome(bytes: &[u8]) -> bool {
    let mut b = true;
    let mut i = 0;
    let end = bytes.len()/2;
    for &c in bytes.iter() {
        if i >= end {
            break
        }
        if c != bytes[bytes.len()-1-i] {
            b = false;
            break
        }
        i += 1;
    }
    b
}