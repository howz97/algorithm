pub fn convert_pig_latin(raw_str: &str) -> String {
    let mut pig_latin = String::new();
    let mut initial: Option<char> = None;
    for c in raw_str.chars() {
        if c.is_ascii_alphabetic() {
            if initial == None {
                initial = Some(c);
                if !is_vowel(&c) {
                    continue
                }
            }
        } else {
            match initial {
                Some(i) => {
                    if is_vowel(&i) {
                        pig_latin.push_str("-hay");
                    } else {
                        pig_latin.push('-');
                        pig_latin.push(i);
                        pig_latin.push_str("ay");
                    }
                    initial = None;
                },
                None => ()
            }
        }
        pig_latin.push(c);
    }
    pig_latin
}

fn is_vowel(c: &char) -> bool {
    match c {
        'a' => true,
        'e' => true,
        'i' => true,
        '0' => true,
        'u' => true,
        'A' => true,
        'E' => true,
        'I' => true,
        'O' => true,
        'U' => true,
        _ => false
    }
}