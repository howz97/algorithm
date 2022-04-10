pub fn my_atoi(s: String) -> i32 {
    let mut v: i64 = 0;
    let mut negative = false;
    for (i, c) in s.trim_start().chars().enumerate() {
       if c.is_ascii_digit() {
           v = v*10 + (c as i64 - '0' as i64);
           if v >= 1<<31 {
                break
           }
       } else if i == 0 && (c == '+' || c == '-') {
            if c == '-' {
                negative = true
            }
       } else {
           break
       }
    }
    if negative {
        if v > 1<<31 {
            v = 1<<31
        }
        v = -v;
    }else if v > (1<<31) - 1 {
        v = (1<<31) - 1
    }
    v as i32
}