use alg_rust::atoi;

#[test]
fn test_atoi() {
    let v = atoi::my_atoi(String::from("9223372036854775808"));
    assert_eq!(v, 2147483647);
}