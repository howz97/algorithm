use alg_rust::median;

#[test]
fn test_find_median_sorted_arrays() {
    let nums1 = vec![1,2];
    let nums2 = vec![3,4];
    let median = median::find_median_sorted_arrays(nums1, nums2);
    assert_eq!(median, 2.5);
}