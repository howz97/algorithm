use alg_rust::palindrome;

#[test]
fn test_longest_palindrome() {
    let s = String::from("abbabdd");
    let lp = palindrome::longest_palindrome(s);
    assert_eq!(lp, "abba");

    let s = String::from("bb");
    let lp = palindrome::longest_palindrome(s);
    assert_eq!(lp, "bb");
}