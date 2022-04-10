pub fn fibonacci(n: i32) -> i32 {
    if n == 0 || n == 1 {
        return 1
    }
    let mut num0 = 1;
    let mut num1 = 1;
    let mut tmp;
    let mut idx_num1 = 1;
    while idx_num1 < n {
        idx_num1 += 1;
        tmp = (num1 + num0) % 1000000007;
        num0 = num1;
        num1 = tmp;
    }
    return num1
}