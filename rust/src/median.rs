pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
    let len1 = nums1.len();
    if len1 == 0 {
        return find_median_sorted_array(&nums2)
    }
    let len2 = nums2.len();
    if len2 == 0 {
        return find_median_sorted_array(&nums1)
    }
    // both nums1 and nums2 contains elements
    let len_sum = len1 + len2;
    let mut next_pop1: usize = 0;
    let mut next_pop2: usize = 0;
    let mut last_pop: (usize, usize) = (0, 0);
    let mut last_last_pop: (usize, usize) = (0, 0);
    let mut need_pop: usize = len_sum/2 + 1;
    while need_pop > 0 {
        let num1: i32;
        match nums1.get(next_pop1) {
            Some(i) => num1 = *i,
            None => {
                // println!("need_pop={}, last_last_pop=({},{}), last_pop=({},{})", need_pop, last_last_pop.0, last_last_pop.1, last_pop.0, last_pop.1);
                pop_n_from_array(&mut last_last_pop, &mut last_pop, need_pop, (2, next_pop2));
                break
            }
        }
        let num2: i32;
        match nums2.get(next_pop2) {
            Some(i) => num2 = *i,
            None => {
                // println!("need_pop={}, last_last_pop=({},{}), last_pop=({},{})", need_pop, last_last_pop.0, last_last_pop.1, last_pop.0, last_pop.1);
                pop_n_from_array(&mut last_last_pop, &mut last_pop, need_pop, (1, next_pop1));
                break
            }
        }
        last_last_pop = last_pop;
        if num1 < num2 {
            last_pop = (1,next_pop1);
            next_pop1 += 1;
        } else {
            last_pop = (2,next_pop2);
            next_pop2 += 1;
        }
        need_pop -= 1;
        // println!("need_pop={}, last_last_pop=({},{}), last_pop=({},{})", need_pop, last_last_pop.0, last_last_pop.1, last_pop.0, last_pop.1);
    }
    let m1: f64 = get_from_2_arrays(last_pop, &nums1, &nums2) as f64;
    if len_sum%2 == 1 {
        // println!("len_sum={}, last_last_pop=({},{}), last_pop=({},{})", len_sum, last_last_pop.0, last_last_pop.1, last_pop.0, last_pop.1);
        println!("m1={}", m1);
        return m1 
    }
    let m0: f64 = get_from_2_arrays(last_last_pop, &nums1, &nums2) as f64;
    // println!("len_sum={}, last_last_pop=({},{}), last_pop=({},{})", len_sum, last_last_pop.0, last_last_pop.1, last_pop.0, last_pop.1);
    // println!("m0={}, m1={}", m0, m1);
    return (m0 + m1)/2.0
}

fn pop_n_from_array(last_last_pop: &mut (usize, usize), last_pop: &mut (usize, usize),
                        need_pop: usize, next_pop: (usize, usize)) {
    if need_pop == 1 {
        *last_last_pop = *last_pop;
        *last_pop = next_pop;
        return 
    }
    *last_last_pop = (next_pop.0, next_pop.1+need_pop-2);
    *last_pop = (next_pop.0, next_pop.1+need_pop-1);
}

fn get_from_2_arrays(idx: (usize, usize), nums1: &Vec<i32>, nums2: &Vec<i32>) -> i32 {
    if idx.0 == 1 {
        if let Some(v) = nums1.get(idx.1) {
            return *v
        } else {
            return 0
        }
    } else {
        if let Some(v) = nums2.get(idx.1) {
            return *v
        } else {
            return 0
        }
    }
}

pub fn find_median_sorted_array(nums: &Vec<i32>) -> f64 {
    let length = nums.len();
    let m: f64;
    match nums.get(length/2) {
        Some(i) => m = *i as f64,
        None => panic!("num not exist")
    }
    if length%2 == 1 {
        return m
    }
    let m0: f64;
    match nums.get(length/2-1) {
        Some(i) => m0 = *i as f64,
        None => panic!("num not exist")
    }
    return (m + m0) / 2.0
}

pub fn calc_median(numbers: &mut Vec<i32>) -> i32 {
    let length = numbers.len();
    if length == 0 {
        panic!("empty vector")
    }
    numbers.sort();
    let mut m = numbers[length / 2];
    if length % 2 == 0 {
        m = (m + numbers[length/2-1])/2
    }
    m
}